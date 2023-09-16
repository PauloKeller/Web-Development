package main

import (
	"log"
	"net"
	"os"

	"hub_service/handler"
	"hub_service/interactors"
	"hub_service/repositories"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	core_grpc_api "hub_service/proto/out"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbdriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	serverPort := ":" + os.Getenv("SERVER_PORT")

	dbData := &repositories.PostgresConnectionData{
		DBName:   dbName,
		Port:     dbPort,
		Password: dbPassword,
		Username: dbUser,
		Driver:   dbdriver,
		Host:     dbHost,
	}

	repositories, err := repositories.NewRepositories(dbData)

	if err != nil {
		panic(err)
	}

	repositories.Automigrate()

	handler := &handler.Handler{
		CreateUserInteractor:  interactors.NewCreateUserInteractor(repositories.User),
		GetAllUserInteractor:  interactors.NewGetAllUserInteractor(repositories.User),
		GetUserByIDInteractor: interactors.NewGetUserByIDInteractor(repositories.User),
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	core_grpc_api.RegisterHubServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	log.Println("Starting server on port", serverPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
