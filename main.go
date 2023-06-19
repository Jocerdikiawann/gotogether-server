package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jocerdikiawann/server_share_trip/config"
	"github.com/Jocerdikiawann/server_share_trip/di"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/Jocerdikiawann/shared_proto_share_trip/auth"
	"github.com/Jocerdikiawann/shared_proto_share_trip/route"
	"github.com/joho/godotenv"
	"github.com/tarndt/wasmws"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tokenDuration = 15 * time.Minute
)

func init() {
	err := godotenv.Load()
	utils.CheckError(err)
}

func main() {
	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	router := http.NewServeMux()
	wsl := wasmws.NewWebSocketListener(appCtx)
	router.HandleFunc("/grpc-proxy", wsl.ServeHTTP)
	httpServer := &http.Server{Addr: ":8888", Handler: router}
	go func() {
		defer appCancel()
		log.Printf("ERROR: HTTP Listen and Server failed; Details: %s", httpServer.ListenAndServe())
	}()

	conf := &config.Config{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		NameDb:   os.Getenv("MONGO_DB_NAME"),
	}

	interceptor := di.InitializedAuthInterceptors(
		conf,
		os.Getenv("SECRET_KEY"),
		tokenDuration,
	)

	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Failed to contruct gRPC TSL credentials from {cert,key}.pem: %s", err)
	}
	serv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
		grpc.Creds(creds),
	)

	routeService := di.InitializedRouteServiceServer(conf)
	authService := di.InitializedAuthServiceServer(
		conf,
		os.Getenv("SECRET_KEY"),
		tokenDuration,
	)

	route.RegisterRouteServer(serv, routeService)
	auth.RegisterAuthServer(serv, authService)

	go func() {
		defer appCancel()
		if err := serv.Serve(wsl); err != nil {
			log.Printf("ERROR: Failed to serve gRPC connections; Details: %s", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Printf("INFO: Received shutdown signal: %s", <-sigs)
		appCancel()
	}()

	<-appCtx.Done()
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*2)
	defer shutdownCancel()

	grpcShutdown := make(chan struct{}, 1)
	go func() {
		serv.GracefulStop()
		grpcShutdown <- struct{}{}
	}()

	httpServer.Shutdown(shutdownCtx)
	select {
	case <-grpcShutdown:
	case <-shutdownCtx.Done():
		serv.Stop()
	}
}
