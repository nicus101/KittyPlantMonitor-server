package server

import "net/http"

type DumpHttpPath string

const defaultLogRotation = 1000

func (path DumpHttpPath) wrapAround( /* TODO */ ) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
