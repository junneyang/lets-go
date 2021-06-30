package main

import (
	"context"
	"fmt"
	"grpc/hello"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

type Server struct {
	hello.UnimplementedGreeterServer
}

func (server *Server) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &hello.HelloReply{Message: fmt.Sprintf("Hello, %v", req.GetName())}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &Server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
