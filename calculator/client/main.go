package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"

	pb "github.com/bgunay/grpc-go-app/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	//doSum(c)
	//fmt.Println(strings.Repeat("-", 22))
	//doPrimes(c)
	//fmt.Println(strings.Repeat("-", 22))
	//doAvg(c)
	fmt.Println(strings.Repeat("-", 22))
	doMax(c)
	//fmt.Println(strings.Repeat("-", 22))
	//doSqrt(c, 10)
	//fmt.Println(strings.Repeat("-", 22))
	//doSqrt(c, -2)
}
