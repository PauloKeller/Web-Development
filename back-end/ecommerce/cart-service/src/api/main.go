package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	cart_grpc "cart-service/grpc"
	"cart-service/handlers"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cs := &handlers.Handlers{}

	cart_grpc.RegisterCartServiceServer(s, cs)

	reflection.Register(s)

	log.Println("Starting Server.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
