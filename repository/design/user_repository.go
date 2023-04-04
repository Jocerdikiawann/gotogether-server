package design

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
)

type AuthRepository interface {
	Authentication(context.Context, request.UserRequest) (entity.Auth, error)
}
