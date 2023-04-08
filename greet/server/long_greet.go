package main

import (
	"fmt"
	"io"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

func (*Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	fmt.Println("Long greet was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		} else if err != nil {
			fmt.Println("streaming error ", err)
		}

		res += fmt.Sprintln("Hello ", req.FirstName)
	}
}
