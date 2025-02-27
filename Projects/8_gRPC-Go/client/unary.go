package main

import (
	"context"
	"log"
	"time"

	pb "github.com/devilzs1/grpc-go/proto"
)

func callSayHello(client pb.GreetServiceClient) () {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Message: %s", resp.Message)

	

}