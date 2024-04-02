package main

import pb "github.com/bgunay/grpc-go-app/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
