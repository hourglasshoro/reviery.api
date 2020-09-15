package main

import (
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

	reflection.Register(server)
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
