package request

import "github.com/Jocerdikiawann/server_share_trip/model/entity"

type DestinationAndPolylineRequest struct {
	GoogleId     string       `validate:"required"`
	Destination  entity.Point `validate:"required"`
	EncodedRoute string       `validate:"required"`
}
