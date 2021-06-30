package main

import (
	"log"
	"net"

	"example.com/rpc-demo/server"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("serving rpc server on ", lis.Addr())
	s := server.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
