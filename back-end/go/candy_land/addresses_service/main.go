package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"candy_land/addresses_service/src/handlers"
	"candy_land/addresses_service/src/repositories"
	"candy_land/addresses_service/src/usecases"
	addresses_grpc "candy_land/grpc"
)

func main() {
	s := grpc.NewServer()

	reflection.Register(s)

	repo := repositories.NewPostgresAddressesRepository()

	h := &handlers.Handler{
		CreateAddressUsecase: &usecases.CreateAddressProvider{
			Repository: repo,
		},
	}

	addresses_grpc.RegisterAddressesServiceServer(s, h)

	port := ":5100"

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
