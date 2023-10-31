package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/nicus101/KittyPlantMonitor-server/pkg/measure"
	"github.com/nicus101/KittyPlantMonitor-server/pkg/sensor"
)

var (
	sensorFindBySerialCodeFn = sensor.FindBySerialCode
)

func SensorIngestV1(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sensorCode := req.Header.Get("Authorization")
	sensorId, err := sensorFindBySerialCodeFn(
		req.Context(),
		sensorCode,
	)
	switch {

	case errors.Is(err, sensor.ErrSensorNotFound):
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return

	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Ingest", sensorId, sensorCode)
	command := struct {
		Api      string            `json:"api"`
		Measures []json.RawMessage `json:"measures"`
	}{}
	err = json.NewDecoder(req.Body).Decode(&command)
	switch {
	case err != nil:
		http.Error(w, "invalid json", http.StatusBadRequest)
		return

	case command.Api != "v0":
		http.Error(w, "only api v0 supported", http.StatusBadRequest)
		return

	case len(command.Measures) == 0:
		http.Error(w, "no measures in command", http.StatusBadRequest)
		return
	}

	values := make([]float64, len(command.Measures))
	for i, measureRaw := range command.Measures {
		values[i] = parseV0Measure(measureRaw)
	}
	err = measure.Insert(req.Context(), sensorId, values)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewSensorMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/sensor/ingest", SensorIngestV1)
	return mux
}

func parseV0Measure(measureRaw json.RawMessage) float64 {
	shim := struct {
		Value float64 `json:"value"`
	}{}
	// TODO: support types
	json.Unmarshal(measureRaw, &shim)
	return shim.Value
}
