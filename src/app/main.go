package main

import (
	pb "github.com/hourglasshoro/reviery.api/src/app/presentation/grpc/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()

	ctlrs, err := NewControllers()
	if err != nil {
		log.Fatalln(err)
	}

	pb.RegisterCommonServer(server, &ctlrs.Common)

	reflection.Register(server)
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
