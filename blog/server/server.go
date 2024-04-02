package main

import pb "github.com/bgunay/grpc-go-app/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
