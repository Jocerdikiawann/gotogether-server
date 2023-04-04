package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/Jocerdikiawann/server_share_trip/config"
	"github.com/Jocerdikiawann/server_share_trip/model/proto/route"
	"github.com/Jocerdikiawann/server_share_trip/repository"
	services "github.com/Jocerdikiawann/server_share_trip/services"
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

	nameDB := os.Getenv("MONGO_DB_NAME")
	portDB := os.Getenv("MONGO_PORT")
	usernameDB := os.Getenv("MONGO_USERNAME")
	passwordDB := os.Getenv("MONGO_PASSWORD")
	hostDB := os.Getenv("MONGO_HOST")

	db := config.Connect(usernameDB, passwordDB, nameDB, hostDB, portDB)

	serv := grpc.NewServer()
	routeRepo := repository.NewRouteRepository(db)
	routeService := &services.RouteServiceServer{
		Repo: routeRepo,
	}
	route.RegisterRouteServer(serv, routeService)

	fmt.Printf("server listening on : %v", listener.Addr())

	err = serv.Serve(listener)
	utils.CheckError(err)
}
