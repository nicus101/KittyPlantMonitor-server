package server

import (
	"context"
	"fmt"
	"net/http"
)

func ListenAndServe() error {
	sensorHandler := http.HandlerFunc(SensorIngestV1)

	apiMux, err := NewDumpMiddleware(
		context.TODO(),
		"log/ingest",
		sensorHandler,
	)
	if err != nil {
		return fmt.Errorf("ingest: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", apiMux)

	return http.ListenAndServe(":6969", mux)
}
