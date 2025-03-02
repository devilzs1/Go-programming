package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/devilzs1/grpc-go/proto"
)


func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Println("Bidirectional streaming started.....")

	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error in stream : %v", err)
	}

	waitch := make(chan struct{})

	go func(){
		for {
			msg, err := stream.Recv()
			if err == io.EOF{
				break;
			}
			if err != nil {
				log.Fatalf("Error while receiving message : %v", err)
			}
			log.Printf("Received message : %v", msg)
		}
		close(waitch)
	}()



	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending req : %v", err)
		}
		time.Sleep(2*time.Second)
	}

	stream.CloseSend()
	<- waitch

	log.Println("Bidirectional Streamming finished!")
}