package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bgunay/grpc-go-app/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var greetWithDeadlineTime = 1 * time.Second

func (*Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(greetWithDeadlineTime)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
