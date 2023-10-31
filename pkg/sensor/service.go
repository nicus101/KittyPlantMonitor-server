package sensor

import (
	"context"
	"errors"
)

var ErrSensorNotFound = errors.New("no such sensor")

func FindBySerialCode(ctx context.Context, serialCode string) (int, error) {
	return selectWhereSerialCode(ctx, serialCode)
}
