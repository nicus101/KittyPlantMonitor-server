package http

import "net/http"

func ListenAndServe() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/plantbed", InspektPlantBed)
	http.HandleFunc("/admin/users", ListUsers)

	authMiddleware := &authMiddleware{
		handler: http.DefaultServeMux,
	}

	return http.ListenAndServe(":6969", authMiddleware)
}
