package main

import (
	"log"

	"github.com/nicus101/KittyPlantMonitor-server/pkg/http"
)

func main() {
	if err := http.ListenAndServe(); err != nil {
		log.Fatal("Serving http failed!", err)
	}
}
