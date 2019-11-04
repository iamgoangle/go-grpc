package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "github.com/iamgoangle/go-grpc/02_server_stream/proto"
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

	grpcClient := pb.NewGreetServiceClient(conn)
	doStreamServer(grpcClient)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

}

func doStreamServer(c pb.GreetServiceClient) {
	req := &pb.CreateRequest{
		Greeting: &pb.Greeting{
			Firstname: "teerapong",
			Lastname:  "singthong",
		},
	}

	resStream, err := c.GreetStream(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling stream RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("Response from stream: %v", msg.GetResult())
	}
}
