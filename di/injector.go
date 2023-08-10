//go:build wireinject
// +build wireinject

package di

import (
	"os"
	"time"

	"github.com/Jocerdikiawann/server_share_trip/db"
	"github.com/Jocerdikiawann/server_share_trip/server"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func provideDBconfig() *db.Config {
	return &db.Config{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		NameDb:   os.Getenv("MONGO_DB_NAME"),
	}
}

func provideSecretKey() string {
	return os.Getenv("SECRET_KEY")
}

func InitializedRouteServiceServer(
	tokenDuration time.Duration,
) *server.RouteServiceServer {
	wire.Build(
		db.MongoDB, validator.New, server.NewRouteRepository, server.NewUserRepository,
		server.NewRouteService, utils.NewLogger, utils.NewJWTManager, provideDBconfig, provideSecretKey,
	)
	return nil
}

func InitializedAuthServiceServer(
	tokenDuration time.Duration,
) *server.UserServiceServer {
	wire.Build(
		db.MongoDB, validator.New,
		server.NewUserRepository,
		server.NewUserService, utils.NewJWTManager,
		utils.NewLogger, provideDBconfig, provideSecretKey,
	)
	return nil
}
