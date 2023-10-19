package http

import (
	"net/http"

	"github.com/nicus101/KittyPlantMonitor-server/pkg/auth"
)

type authMiddleware struct {
	handler http.Handler
}

func (a *authMiddleware) ServeHTTP(
	w http.ResponseWriter, r *http.Request,
) {
	if login, password, valid := r.BasicAuth(); valid {
		ctx, err := auth.LogIn(r.Context(), login, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r = r.WithContext(ctx)
	}

	a.handler.ServeHTTP(w, r)
}
