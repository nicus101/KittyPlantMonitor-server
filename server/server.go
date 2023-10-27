package server

import "net/http"

func ListenAndServe() error {
	mux := http.NewServeMux()

	return http.ListenAndServe(":6969", mux)
}
