package main

import (
	"context"
	"grpc_practice/ping"
	pb "grpc_practice/protobuf/ping"
	"log"
	"time"

	"google.golang.org/grpc"
)

const address = ":8080"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingPongServiceClient(conn)

	// Contact the server and print out its response.
	str := "LeviTest"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PingPong(ctx, ping.PingSomething(str))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("GetPong: %s", r.GetPong())
}
