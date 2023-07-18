package design

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
)

type RouteRepository interface {
	GetDestinationAndPolyline(context context.Context, id string) (entity.Destination, error)
	SendDestinationAndPolyline(context context.Context, request request.DestinationAndPolylineRequest) (string, error)
}
