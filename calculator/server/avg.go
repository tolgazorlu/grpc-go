package main

import (
	"io"
	"log"

	pb "github.com/tolgazorlu/grpc-go/calculator/proto"
)

func (*Server) AVG(stream pb.CalculatorService_AVGServer) error {
	log.Println("AVG was inwoked")

	var sum int32 = 0
	count := 0

	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AVGResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		sum += req.Value
		count++
	}
}
