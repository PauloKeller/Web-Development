package main

import (
	"log"
	"net"

	"candy_land/orders_service/src/handlers"
	"candy_land/orders_service/src/repositories"
	"candy_land/orders_service/src/usecases"

	candyland_grpc "candy_land/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	reflection.Register(s)

	repo := repositories.NewPostgresOrdersRepository()

	h := &handlers.Handler{
		CreateOrderUsecase: &usecases.CreateOrderProvider{
			Repository: repo,
		},
	}

	candyland_grpc.RegisterOrdersServiceServer(s, h)

	port := ":5300"

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
