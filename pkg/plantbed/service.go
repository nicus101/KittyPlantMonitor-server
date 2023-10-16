package plantbed

import "context"

func ById(ctx context.Context, plantBedId int) (Response, error) {

	return Response{Humidity: 13}, nil
}
