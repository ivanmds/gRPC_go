package main

import (
	"context"
	"errors"

	pb "github.com/ivanmds/gRPC_go/calc/proto"
)

func (*Server) Sum(ctx context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {
	if request == nil {
		return nil, errors.New("should be informed a values")
	}

	var sunResult = request.Value1 + request.Value2

	return &pb.SumResponse{Result: sunResult}, nil
}
