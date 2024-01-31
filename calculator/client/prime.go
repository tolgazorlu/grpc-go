package main

import (
	"context"
	"io"
	"log"

	pb "github.com/tolgazorlu/grpc-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimeRequest{
		Value: 12390392840,
	}
	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res.Result)
	}
}
