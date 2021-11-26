package main

import (
	"candy_land/carts_service/src/handlers"

	"candy_land/carts_service/src/repositories"
	"candy_land/carts_service/src/usecases"
	"log"
	"net"

	candyland_grpc "candy_land/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	reflection.Register(s)

	repo := repositories.NewPostgresCartsRepository()

	h := &handlers.Handler{
		AddProductToCartUsecase: &usecases.AddProductToCartProvider{
			Repository: repo,
		},
	}

	candyland_grpc.RegisterCartsServiceServer(s, h)

	port := ":5200"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	log.Println("Starting server on port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
