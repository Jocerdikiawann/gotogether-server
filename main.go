package main

import (
	"context"
	"flag"
	"fmt"
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
	tokenDuration      = time.Hour * 24 * 365
)

func init() {
	err := godotenv.Load()
	utils.CheckError(err)
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcMux := runtime.NewServeMux()

	conf := &config.Config{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		NameDb:   os.Getenv("MONGO_DB_NAME"),
	}

	routeService := di.InitializedRouteServiceServer(
		conf,
		os.Getenv("SECRET_KEY"),
		tokenDuration,
	)
	authService := di.InitializedAuthServiceServer(
		conf,
		os.Getenv("SECRET_KEY"),
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

	fmt.Printf("Server run on http://%v", s.Addr)

	if servError := s.ListenAndServe(); servError != nil {
		utils.CheckError(servError)
	}
}