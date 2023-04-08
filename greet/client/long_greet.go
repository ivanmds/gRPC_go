package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

func doLognGreet(client pb.GreetServiceClient) {

	reqs := []*pb.GreetRequest{
		{FirstName: "Test1"},
		{FirstName: "Test2"},
		{FirstName: "Test3"},
		{FirstName: "Test4"},
		{FirstName: "Test5"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		fmt.Println("streaming error ", err)
	}

	for _, req := range reqs {
		fmt.Println("Sending ", req.FirstName)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		fmt.Println("Error while receiving response")
	}

	fmt.Println("Long greet said\n", res.Result)
}
