package main

import (
	"context"
	"fmt"

	pb "github.com/ivanmds/gRPC_go/calc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("ConnectionClient error", err)
	}
	defer conn.Close()

	client := pb.NewSumServiceClient(conn)
	doCalc(client)
}

func doCalc(client pb.SumServiceClient) {
	var value1 uint32 = 1
	var value2 uint32 = 2
	for {
		if value1 >= 1000000 {
			break
		}

		sunRequest := &pb.SumRequest{Value1: value1, Value2: value2}

		res, err := client.Sum(context.Background(), sunRequest)
		if err != nil {
			fmt.Println("error when call server ", err)
		}

		fmt.Println("the result from sum value1 ", sunRequest.Value1, " and value2 ", sunRequest.Value2, " is ", res.Result)

		value1++
		value2 = value2 + value1
	}
}
