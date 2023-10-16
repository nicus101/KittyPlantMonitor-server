package http

import (
	"net/http"
	"text/template"

	"github.com/nicus101/KittyPlantMonitor-server/pkg/plantbed"
)

func InspektPlantBed(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("plant_bed.html").ParseFiles("pkg/http/views/plant_bed.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, plantbed.Response{
		Humidity: 69,
	})
	if err != nil {
		panic(err)
	}
}
