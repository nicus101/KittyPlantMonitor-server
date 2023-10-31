package measure

import (
	"context"
	"database/sql"
	"strings"
)

var dbConn *sql.DB

func SetSqlDb(db *sql.DB) {
	dbConn = db
}

func insertMeasures(
	ctx context.Context, sensorId int, measures []float64,
) error {
	args := make([]any, len(measures)*2)
	values := make([]string, len(measures))
	for i, measure := range measures {
		args[i*2] = sensorId
		args[i*2+1] = measure
		values[i] = `(?,?)`
	}

	query := "INSERT INTO measures(sensor_id, value) VALUES" + strings.Join(values, ",")
	//fmt.Println(query)
	_, err := dbConn.ExecContext(
		ctx,
		query,
		args...,
	)

	return err
}
