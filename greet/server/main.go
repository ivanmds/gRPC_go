package main

import (
	"log"
	"net"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf("Failed to listen on: %v\n", err)
		return
	}
	defer lis.Close()

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	defer s.Stop()

	if err = s.Serve(lis); err != nil {
		log.Printf("Failed to server: %v\n", err)
		return
	}
}
