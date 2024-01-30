package main

import (
	"context"
	"log"

	pb "github.com/tolgazorlu/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstValue:  2,
		SecondValue: 4,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)

}
