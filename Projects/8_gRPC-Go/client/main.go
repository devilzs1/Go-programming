package main

import (
	"fmt"
	"log"

	pb "github.com/devilzs1/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HelloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	fmt.Println("Client")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// conn, err := grpc.DialContext(ctx, "localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("dns:///localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"CodeForMe", "codeforus", "devilzs1"},
	}
	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
