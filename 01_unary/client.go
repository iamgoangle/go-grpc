package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "github.com/iamgoangle/go-grpc/01_unary/proto"
)

const (
	address = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect grpc server %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := grpcClient.Create(ctx, &pb.CreateRequest{
		Api: "v1",
		Todo: &pb.ToDo{
			Id:         1,
			Title:      "Golf test",
			Descripton: "Test grpc for create new todo",
			Reminder:   ptypes.TimestampNow(),
		},
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %+v", r.GetTodo())
}
