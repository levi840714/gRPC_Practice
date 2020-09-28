package ping

import (
	"context"
	pb "grpc_practice/protobuf/ping"
	"log"
)

type PingPongService struct {
}

func (p *PingPongService) PingPong(ctx context.Context, req *pb.PingRequest) (*pb.PongResponse, error) {
	log.Printf("GetPing: %s", req.GetPing())
	return &pb.PongResponse{Pong: req.GetPing()}, nil
}
