package request

import "github.com/Jocerdikiawann/server_share_trip/model/entity"

type LocationRequest struct {
	GoogleId string       `json:"googleId" validate:"required"`
	Point    entity.Point `json:"point" validate:"required"`
	IsFinish bool         `json:"isFinish" validate:"required"`
}
