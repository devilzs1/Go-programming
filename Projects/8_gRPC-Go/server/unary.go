package main

import (
	"context"
	pb "github.com/devilzs1/grpc-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello world!",
	}, nil
}