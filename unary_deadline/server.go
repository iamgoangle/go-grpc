package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/iamgoangle/go-grpc/unary_deadline/proto"
)

const (
	port = ":9090"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Started gRPC server port: %v\n", port)
}

func (s *server) HelloWithDeadline(ctx context.Context, req *pb.HelloWithDeadlineRequest) (*pb.HelloWithDealLineResponse, error) {
	fmt.Printf("HelloWithDeadline function was invoked with %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetHello().GetFirstName()
	result := "Hello " + firstName
	res := &pb.HelloWithDealLineResponse{
		Result: result,
	}

	return res, nil
}
