package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

func (*Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	if in == nil {
		log.Println("request is empty")
		return nil, errors.New("request is empty")
	}

	log.Printf("Greet funcion was invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
