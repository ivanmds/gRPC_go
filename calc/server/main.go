package main

import (
	"fmt"
	"net"

	pb "github.com/ivanmds/gRPC_go/calc/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("Listening error ", err)
	}

	defer lis.Close()
	fmt.Println("Listening at ", addr)

	server := grpc.NewServer()
	pb.RegisterSumServiceServer(server, &Server{})

	defer server.Stop()

	if err = server.Serve(lis); err != nil {
		fmt.Println("Server error ", err)
	}
}
