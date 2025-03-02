package main

import (
	"context"
	"log"
	"time"

	pb "github.com/devilzs1/grpc-go/proto"
)


func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client streaming started.....")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not stream : %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Could not send names : %v", err)
		}

		log.Printf("Send the request with name : %v", name)
		time.Sleep(2*time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving : %v", err)
	}
	log.Printf("Message : %v", res.Messages)
	log.Println("Client streaming finished.....")
}