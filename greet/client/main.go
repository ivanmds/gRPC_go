package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ivanmds/gRPC_go/greet/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	//doGreet(c)
	//doGreetManyTimes(c)
	//doLognGreet(c)
	doGreetEveryone(c)
}
