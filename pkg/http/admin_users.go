package http

import (
	"fmt"
	"net/http"

	"github.com/nicus101/KittyPlantMonitor-server/pkg/auth"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	if !requireAdmin(w, r) {
		return
	}

	fmt.Fprintln(w, "yaaay")
}

func requireAdmin(w http.ResponseWriter, r *http.Request) bool {
	user, valid := auth.CtxUser(r.Context())
	if !valid {
		w.Header().Set("WWW-Authenticate", `Basic realm="panel", charset="UTF-8"`)
		http.Error(w, "Miau Miau", http.StatusUnauthorized)
		return false
	}
	if !user.IsAdmin() {
		http.Error(w, "Permission denied", http.StatusForbidden)
		return false
	}
	return true
}
