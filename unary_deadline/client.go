package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	pb "github.com/iamgoangle/go-grpc/unary_deadline/proto"
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

	grpcClient := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := grpcClient.HelloWithDeadline(ctx, &pb.HelloWithDeadlineRequest{
		Hello: &pb.Hello{
			FirstName: "golf",
			LastName:  "singthong",
		},
	})

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("timeout na ja")
			} else {
				log.Println(err)
			}
		} else {
			log.Fatalf("could not greet: %v", err)
		}

	}

	log.Printf("Response: %+v", r.GetResult())
}
