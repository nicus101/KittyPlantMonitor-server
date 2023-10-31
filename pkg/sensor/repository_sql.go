package sensor

import (
	"context"
	"database/sql"
)

var dbConn *sql.DB

func SetSqlDb(db *sql.DB) {
	dbConn = db
}

func selectWhereSerialCode(
	ctx context.Context, serialCode string,
) (
	int, error,
) {
	row := dbConn.QueryRowContext(
		ctx,
		"SELECT id FROM sensors WHERE serial_code = ?",
		serialCode,
	)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int
	row.Scan(&id)
	if id == 0 {
		return 0, ErrSensorNotFound
	}

	return id, nil
}
