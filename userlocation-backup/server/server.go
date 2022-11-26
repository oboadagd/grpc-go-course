package main

import pb "github.com/Clement-Jean/grpc-go-course/userlocation/proto"

type Server struct {
	pb.UserLocationServiceServer
}
