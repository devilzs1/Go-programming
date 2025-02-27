package main

import (
	"log"
	"time"

	pb "github.com/devilzs1/grpc-go/proto"
)

func (s *helloServer) callSayHelloServerStream(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error{

	log.Printf("Request with names : %v", req.Names)
	for _, name := range req.Names {
		resp := &pb.HelloResponse{
			Message: "Hello " + name + "!",
		}
		if err := stream. Send(resp); err != nil {
			log.Fatalf("Error sending response from server: %v", err)
			return err
		}
	}
	time.Sleep(1*time.Second)

	return nil
}