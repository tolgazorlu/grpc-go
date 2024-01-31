package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tolgazorlu/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("longGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Tolga"},
		{FirstName: "Mustafa"},
		{FirstName: "Ayşe"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
