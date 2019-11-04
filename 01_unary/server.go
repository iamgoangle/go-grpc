package main

import (
	pb "github.com/iamgoangle/go-grpc/01_unary/proto"
)

const (
	port = ":9090"
)

type server struct {
}

func main() { 
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}