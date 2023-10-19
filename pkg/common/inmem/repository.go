package inmem

import (
	"encoding/json"
	"log"
)

type table []map[string]any

var tables = map[string]table{}

func LoadFromJsonBytes(name string, data []byte) {
	var table table
	err := json.Unmarshal(data, &table)
	if err != nil {
		log.Fatalln("Loading table from json:", err)
	}
	tables[name] = table
}
