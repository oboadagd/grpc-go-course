package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/userlocation/proto"
)

func doPing(c pb.UserLocationServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Ping(context.Background(), &pb.PingRequest{Message: "Pong"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", r.Result)
}
