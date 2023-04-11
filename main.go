package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/Jocerdikiawann/server_share_trip/config"
	"github.com/Jocerdikiawann/server_share_trip/di"
	"github.com/Jocerdikiawann/server_share_trip/model/proto/route"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8888, "server port")
)

func init() {
	err := godotenv.Load()
	utils.CheckError(err)
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	utils.CheckError(err)

	serv := grpc.NewServer()

	conf := &config.Config{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		NameDb:   os.Getenv("MONGO_DB_NAME"),
	}

	routeService := di.InitializedServiceServer(conf)

	route.RegisterRouteServer(serv, routeService)

	fmt.Printf("server listening on : %v", listener.Addr())

	err = serv.Serve(listener)
	utils.CheckError(err)
}
