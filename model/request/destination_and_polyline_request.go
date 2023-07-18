package request

import "github.com/Jocerdikiawann/server_share_trip/model/entity"

type DestinationAndPolylineRequest struct {
	GoogleId     string       `json:"googleId" bson:"googleId,omitempty" validate:"required"`
	Destination  entity.Point `json:"destination" bson:"destination,omitempty" validate:"required"`
	EncodedRoute string       `json:"encodedRoute" bson:"encodedRoute,omitempty" validate:"required"`
}
