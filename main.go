package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Jocerdikiawann/server_share_trip/config"
	"github.com/Jocerdikiawann/server_share_trip/di"
	"github.com/Jocerdikiawann/server_share_trip/model/pb"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8888", "gRPC server endpoint")
	tokenDuration      = 15 * time.Minute
	secretKey          = os.Getenv("SECRET_KEY")
)

func init() {
	err := godotenv.Load()
	utils.CheckError(err)
}

func run() error {
	conf := &config.Db{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		NameDb:   os.Getenv("MONGO_DB_NAME"),
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcMux := runtime.NewServeMux()

	routeService := di.InitializedRouteServiceServer(
		conf,
		secretKey,
		tokenDuration,
	)
	authService := di.InitializedAuthServiceServer(
		conf,
		secretKey,
		tokenDuration,
	)

	routeErr := pb.RegisterRouteHandlerServer(ctx, grpcMux, routeService)
	utils.CheckError(routeErr)
	authErr := pb.RegisterAuthHandlerServer(ctx, grpcMux, authService)
	utils.CheckError(authErr)

	s := &http.Server{
		Addr:    *grpcServerEndpoint,
		Handler: grpcMux,
	}

	return s.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("cannot start GRPC server %v", err)
	}
}
