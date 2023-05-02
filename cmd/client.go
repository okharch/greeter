package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/okharch/greeter"
	"google.golang.org/grpc"
)

func main() {
	// Parse command-line arguments.
	portPtr := flag.Int("port", 50051, "the server port")
	flag.Parse()

	// Create a gRPC connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf(":%d", *portPtr), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	c := pb.NewGreeterClient(conn)

	// Call the gRPC method on the server.
	name := "World"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Say: %s", r.Message)
	r, err = c.SayTime(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not get time: %v", err)
	}
	log.Printf("Time: %s", r.Message)
}
