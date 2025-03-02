package main

import (
	"io"
	"log"

	pb "github.com/devilzs1/grpc-go/proto"
)


func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving req : %v", err)
			return err
		}
		log.Printf("Got request with nae : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Error sending response : %v", err)
			return err
		}

	}
}