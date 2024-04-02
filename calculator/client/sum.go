package main

import (
	"context"
	"log"

	pb "github.com/bgunay/grpc-go-app/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	r, err := c.Sum(context.Background(),
		&pb.SumRequest{
			FirstNumber:  11,
			SecondNumber: 22,
		},
	)

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", r.Result)
}
