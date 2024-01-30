package main

import (
	"context"
	"log"

	pb "github.com/tolgazorlu/grpc-go/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstValue + in.SecondValue,
	}, nil
}
