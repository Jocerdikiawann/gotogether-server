//go:build wireinject
// +build wireinject

package di

import (
	"time"

	"github.com/Jocerdikiawann/server_share_trip/config"
	"github.com/Jocerdikiawann/server_share_trip/interceptors"
	"github.com/Jocerdikiawann/server_share_trip/repository"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/services"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var routeSet = wire.NewSet(
	repository.NewRouteRepository,
	wire.Bind(new(design.RouteRepository), new(*repository.RouteRepositoryImpl)),
)

var authSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(design.AuthRepository), new(*repository.AuthRepositoryImpl)),
)

var teleSet = wire.NewSet(
	repository.NewTelegramRepository,
	wire.Bind(new(design.TelegramRepository), new(*repository.TelegramRepositoryImpl)),
)

func InitializedRouteServiceServer(
	conf *config.Config,
) *services.RouteServiceServer {
	wire.Build(
		config.Connect, routeSet, validator.New, services.NewRouteService,
	)
	return nil
}

func InitializedAuthServiceServer(
	conf *config.Config,
	token string,
	tokenDuration time.Duration,
) *services.UserServiceServer {
	wire.Build(
		config.Connect, authSet, validator.New, services.NewUserService, utils.NewJWTManager,
	)
	return nil
}

func InitializedAuthInterceptors(
	conf *config.Config,
	token string,
	tokenDuration time.Duration,
) *interceptors.AuthInterceptor {
	wire.Build(config.Connect, utils.NewJWTManager, authSet, interceptors.NewAuthInterceptor, utils.NewLogger, teleSet)
	return nil
}
