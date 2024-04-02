package main

import pb "github.com/bgunay/grpc-go-app/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
