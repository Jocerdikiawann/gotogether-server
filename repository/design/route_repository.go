package design

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
)

type RouteRepository interface {
	GetLocation(context context.Context, id string) (entity.Location, error)
	GetDestinationAndPolyline(context context.Context, id string) (entity.Destination, error)
	SendLocation(context context.Context, request request.LocationRequest) error
	SendDestinationAndPolyline(context context.Context, request request.DestinationAndPolylineRequest) error
}
