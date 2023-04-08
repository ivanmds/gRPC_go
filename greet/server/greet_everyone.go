package main

import (
	"fmt"
	"io"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

func (*Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	fmt.Println("greet everyone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		} else if err != nil {
			fmt.Println("server receive streaming error ", err)
		}

		resp := "Hello " + req.FirstName
		fmt.Println(resp)
		err = stream.Send(&pb.GreetResponse{Result: resp})

		if err != nil {
			fmt.Println("Error server response ", err)
		}
	}
}
