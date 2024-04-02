package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/bgunay/grpc-go-app/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change that to true if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	opts = append(opts, grpc.WithChainUnaryInterceptor(LogInterceptor(), AddHeaderInterceptor()))

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	fmt.Println("--------doGreet(c)--------")
	doGreet(c)
	//fmt.Println("--------doGreetManyTimes(c)--------")
	//doGreetManyTimes(c)
	//fmt.Println("--------doLongGreet(c)--------")
	//doLongGreet(c)
	//fmt.Println("--------doGreetEveryone(c)--------")
	//doGreetEveryone(c)
	//doGreetWithDeadline(c, 5*time.Second)
	//doGreetWithDeadline(c, 1*time.Second)
}
