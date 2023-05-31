package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Jocerdikiawann/server_share_trip/config"
	"github.com/Jocerdikiawann/server_share_trip/di"
	"github.com/Jocerdikiawann/server_share_trip/model/proto/auth"
	"github.com/Jocerdikiawann/server_share_trip/model/proto/route"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port          = flag.Int("port", 8888, "server port")
	tokenDuration = 15 * time.Minute
)

func init() {
	err := godotenv.Load()
	utils.CheckError(err)
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	utils.CheckError(err)

	interceptor := di.InitializedAuthInterceptors(os.Getenv("SECRET_KEY"), tokenDuration)

	serv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	conf := &config.Config{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		NameDb:   os.Getenv("MONGO_DB_NAME"),
	}

	routeService := di.InitializedRouteServiceServer(conf)
	authService := di.InitializedAuthServiceServer(
		conf,
		os.Getenv("SECRET_KEY"),
		tokenDuration,
	)

	route.RegisterRouteServer(serv, routeService)
	auth.RegisterAuthServer(serv, authService)

	fmt.Printf("server listening on : %v", listener.Addr())

	err = serv.Serve(listener)
	utils.CheckError(err)
}
