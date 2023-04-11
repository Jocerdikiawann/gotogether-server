//go:build wireinject
// +build wireinject

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

func InitializedServiceServer(
	conf *config.Config,
) *services.RouteServiceServer {
	wire.Build(
		config.Connect, routeSet, services.NewRouteService,
	)
	return nil
}
