package request

import "github.com/Jocerdikiawann/server_share_trip/model/entity"

type DestinationAndPolylineRequest struct {
	GoogleId        string        `json:"googleId" bson:"googleId,omitempty" validate:"required"`
	Destination     *entity.Point `json:"destination" bson:"destination,omitempty" validate:"required"`
	InitialLocation *entity.Point `json:"initialLocation" bson:"initialLocation,omitempty" validate:"required"`
	EncodedRoute    string        `json:"encodedRoute" bson:"encodedRoute,omitempty" validate:"required"`
	LocationName    string        `json:"locationName" bson:"locationName" validate:"required"`
	DestinationName string        `json:"destinationName" bson:"destinationName" validate:"required"`
	EstimateTime    string        `json:"estimateTime" bson:"estimateTime" validate:"required"`
}
