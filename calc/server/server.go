package main

import pb "github.com/ivanmds/gRPC_go/calc/proto"

type Server struct {
	pb.SumServiceServer
}
