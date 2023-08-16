package main

import (
	"context"

	"github.com/rob-bender/go-grpc-test/pb"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello World",
	}, nil
}
