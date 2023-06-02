package design

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
)

type AuthRepository interface {
	SignUp(context.Context, request.UserRequest) (entity.Auth, error)
	CheckIsValidEmail(context.Context, string) (bool, error)
}
