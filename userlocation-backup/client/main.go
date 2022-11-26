package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/userlocation/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50061"

func main() {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserLocationServiceClient(conn)

	doPing(c)
}
