package main

import (
	"log"

	"github.com/nicus101/KittyPlantMonitor-server/server"
)

func main() {
	var err error
	log.Println("Server starting...")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failure!", err)
	}
}
