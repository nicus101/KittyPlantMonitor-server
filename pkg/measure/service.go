package measure

import "context"

func Insert(
	ctx context.Context, sensorId int, measures []float64,
) error {
	return insertMeasures(ctx, sensorId, measures)
}
