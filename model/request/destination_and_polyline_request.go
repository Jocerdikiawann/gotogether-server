package request

import "github.com/Jocerdikiawann/server_share_trip/model/entity"

type DestinationAndPolylineRequest struct {
	GoogleId    string
	Destination entity.Point
	Polyline    []entity.Point
}
