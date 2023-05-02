package main

import (
	"context"
	"log"
	"net"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
)

func main() {
	// Create a gRPC server on port 50051 to forward requests to.
	dst, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer dst.Close()

	// Create a director function that always forwards requests to the same gRPC server.
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		log.Printf("proxying %s", fullMethodName)
		return ctx, dst, nil
	}

	// Create a gRPC proxy on port 50050 that forwards all requests to the server using the director function.
	proxy := grpc.NewServer(grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Starting gRPC proxy on port 50050")
	if err := proxy.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
