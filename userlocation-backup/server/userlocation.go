package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/userlocation/proto"
)

func (*Server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Ping was invoked with %v\n", in)
	return &pb.PingResponse{Result: "Ping -> " + in.Message}, nil
}
