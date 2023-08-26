package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Jocerdikiawann/server_share_trip/di"
	"github.com/Jocerdikiawann/server_share_trip/model/pb"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {
	err := godotenv.Load()
	utils.CheckError(err)
}

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8888", "gRPC server endpoint")
	tokenDuration      = time.Hour * 24 * 365
)

func newServer() *http.Server {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcMux := runtime.NewServeMux()

	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(grpcMux)

	routeService := di.InitializedRouteServiceServer(tokenDuration)
	authService := di.InitializedAuthServiceServer(tokenDuration)

	routeErr := pb.RegisterRouteHandlerServer(ctx, grpcMux, routeService)
	utils.CheckError(routeErr)
	authErr := pb.RegisterAuthHandlerServer(ctx, grpcMux, authService)
	utils.CheckError(authErr)

	s := &http.Server{
		Addr:    *grpcServerEndpoint,
		Handler: withCors,
	}
	return s
}

func main() {
	s := newServer()
	fmt.Printf("Server run on %v", s.Addr)

	if servError := s.ListenAndServe(); servError != nil {
		utils.CheckError(servError)
	}
}
