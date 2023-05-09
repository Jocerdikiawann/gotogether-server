package di

import (
	"github.com/Jocerdikiawann/server_share_trip/config"
	"github.com/Jocerdikiawann/server_share_trip/repository"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/services"
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

func InitializedAuthServiceServer() *services.UserServiceServer {
	
	return nil
}
