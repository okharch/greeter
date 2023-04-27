package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "grpc-example"
)

func main() {
    conn, err := grpc.Dial(":50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)
    name := "World"
    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Message)
}

