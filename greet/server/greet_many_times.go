package main

import (
	"fmt"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

func (*Server) GreetManyTimes(request *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	fmt.Println("greet many times requested with ", request)

	for i := 0; i < 1000000; i++ {
		res := fmt.Sprintln("Hello ", request.FirstName, ", number ", i)
		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
