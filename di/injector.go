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

func InitializedRouteServiceServer(
	conf *config.Config,
) *services.RouteServiceServer {
	wire.Build(
		config.Connect, routeSet, services.NewRouteService,
	)
	return nil
}

func InitializedAuthServiceServer(
	conf *config.Config,
	token string,
	tokenDuration time.Duration,
) *services.UserServiceServer {
	wire.Build(
		config.Connect, authSet, services.NewUserService, utils.NewJWTManager,
	)
	return nil
}

func InitializedAuthInterceptors(
	token string,
	tokenDuration time.Duration,
) *interceptors.AuthInterceptor {
	wire.Build(utils.NewJWTManager, interceptors.NewAuthInterceptor)
	return nil
}
