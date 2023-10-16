package http

import "net/http"

func ListenAndServe() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/plantbed", InspektPlantBed)

	return http.ListenAndServe(":6969", nil)
}
