package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/iamgoangle/go-grpc/02_server_stream/proto"
)

const (
	port = ":9090"
)

type server struct {
}

func (s *server) GreetStream(req *pb.CreateRequest, stream pb.GreetService_GreetStreamServer) error {
	log.Printf("Prepare data to stream response to client: %v", req)

	fullname := req.GetGreeting().GetFirstname()
	for i := 0; i < 10; i++ {
		result := "Hello " + fullname + " of " + strconv.Itoa(i)
		res := &pb.CreateResponse{
			Result: result,
		}

		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Started gRPC server port: %v\n", port)
}
