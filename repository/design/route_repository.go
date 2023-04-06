package design

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteRepository interface {
	WatchLocation() (*mongo.ChangeStream, error)
	GetDestinationAndPolyline(context context.Context, id string) (entity.Destination, error)
	SendLocation(context context.Context, request request.LocationRequest) (string, error)
	SendDestinationAndPolyline(context context.Context, request request.DestinationAndPolylineRequest) (string, error)
}
