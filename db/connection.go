package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nicus101/KittyPlantMonitor-server/pkg/measure"
	"github.com/nicus101/KittyPlantMonitor-server/pkg/sensor"
)

func Connect(dsn string) error {
	db, err := sql.Open("sqlite3", "db/database.sqlite3")
	if err != nil {
		return err
	}

	// TODO: maybe migrate ?

	// dependency inversion - sql.Db is universal enough
	sensor.SetSqlDb(db)
	measure.SetSqlDb(db)
	return nil
}
