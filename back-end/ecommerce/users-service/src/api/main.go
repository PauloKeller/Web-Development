package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	users_grpc "users-service/grpc"
	"users-service/handlers"
	"users-service/repositories"
	"users-service/usecases"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	// Create a customized counter metric.
	// customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
	// 	Name: "demo_server_say_hello_method_handle_count",
	// 	Help: "Total number of RPCs handled on the server.",
	// }, []string{"name"})
)

func main() {
	cacheRepo := repositories.NewRedisUsersRepository()
	repo := repositories.NewPostgresUsersRepository()

	h := &handlers.Handlers{
		GetUserUsecase: &usecases.GetUserImpl{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
		CreateUserUsecase: &usecases.CreateUserImpl{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
	}

	port := ":" + os.Getenv("PORT")
	metricsPort := os.Getenv("METRICS-PORT")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%s", metricsPort)}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	users_grpc.RegisterUserServiceServer(s, h)

	// Initialize all metrics.
	grpcMetrics.InitializeMetrics(s)

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	reflection.Register(s)

	log.Println("Starting server on port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
