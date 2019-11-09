package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"

	pb "github.com/iamgoangle/go-grpc/unary/proto"
)

const (
	port = ":9090"
)

type server struct {
	pb.TodoServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("Create: %v", req.Todo.GetTitle())

	return &pb.CreateResponse{
		Api: "v1",
		Todo: &pb.ToDo{
			Id:         int64(rand.Intn(100)),
			Title:      req.GetTodo().GetTitle(),
			Descripton: req.GetTodo().GetDescripton(),
			Reminder:   req.GetTodo().GetReminder(),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Started gRPC server port: %v\n", port)
}
