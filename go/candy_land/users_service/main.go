package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	candyland_grpc "candy_land/grpc"

	"candy_land/users_service/src/handlers"
	"candy_land/users_service/src/repositories"
	"candy_land/users_service/src/usecases"
)

func main() {
	s := grpc.NewServer()

	reflection.Register(s)

	repo := repositories.NewPostgresUsersRepository()

	h := &handlers.Handler{
		CreateUserUsecase: &usecases.CreateUserProvider{
			Repository: repo,
		},
	}

	candyland_grpc.RegisterUsersServiceServer(s, h)

	port := ":5000"

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
