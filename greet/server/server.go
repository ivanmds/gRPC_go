package main

import pb "github.com/ivanmds/gRPC_go/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
