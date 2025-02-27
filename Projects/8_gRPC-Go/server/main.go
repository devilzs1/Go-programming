package main

import (
	"fmt"
	"log"
	"net"

    pb "github.com/devilzs1/grpc-go/proto"
	"google.golang.org/grpc"
)

type helloServer struct{
    pb.GreetServiceServer
}

const (
    port = ":8080"
)

func main() {
    fmt.Println("Server")
    listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to start gRPC: %v", err)
    }

}