package main

import (
	"context"
	"log"

	pb "github.com/bgunay/grpc-go-app/greet/proto"
)

func (*Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
