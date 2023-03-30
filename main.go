package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Jocerdikiawann/server_share_trip/model"
	"github.com/Jocerdikiawann/server_share_trip/services"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8888, "server port")
)

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := services.NewService()
	model.RegisterShareTripServer(server, service)
	fmt.Printf("server listening on %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
