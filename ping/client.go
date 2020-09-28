package ping

import pb "grpc_practice/protobuf/ping"

func PingSomething(str string) *pb.PingRequest {
	return &pb.PingRequest{Ping: str}
}
