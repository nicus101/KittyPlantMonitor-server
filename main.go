package main

import (
	"log"

	"github.com/nicus101/KittyPlantMonitor-server/db"
	"github.com/nicus101/KittyPlantMonitor-server/server"
)

func main() {
	var err error
	log.Println("Server starting...")

	err = db.Connect("")
	if err != nil {
		log.Fatal("Db connection failure!", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failure!", err)
	}
}
