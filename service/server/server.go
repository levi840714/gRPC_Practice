package main

import (
	"log"
	"net"

	"grpc_practice/ping"
	pb "grpc_practice/protobuf/ping"

	"google.golang.org/grpc"
)

const port = ":8080"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	log.Println("gRPC server is running.")

	p := &ping.PingPongService{}
	pb.RegisterPingPongServiceServer(s, p)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
