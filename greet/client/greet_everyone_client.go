package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

func doGreetEveryone(client pb.GreetServiceClient) {
	fmt.Println("doGreetEveryone was invoked")

	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		fmt.Println("Streaming error ", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Test1"},
		{FirstName: "Test2"},
		{FirstName: "Test3"},
		{FirstName: "Test4"},
		{FirstName: "Test5"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			stream.Send(req)
			time.Sleep(time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("error to receive response stream ", err)
				break
			}

			fmt.Println("Response stream: ", res.Result)
		}

		close(waitc)
	}()
	<-waitc
}
