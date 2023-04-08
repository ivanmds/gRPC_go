package main

import (
	"context"
	"fmt"
	"io"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

func doGreetManyTimes(client pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		FirstName: "Ivan",
	}

	stream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		fmt.Println("Error when called the server ", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("Stream error ", err)
		} else if err != nil {
			break
		}

		fmt.Println(msg.Result)
	}
}
