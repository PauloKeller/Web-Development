package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"products-service/handlers"
	"products-service/repositories"
	"products-service/usecases"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "products_api_request_time",
		Help:       "Time it takes to handle a Request.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	userRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "products_api_requests",
			Help: "Number of requests for the API.",
		},
		[]string{"request"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(requestTime)
	prometheus.MustRegister(userRequests)
}

func main() {
	repo := repositories.NewPostgresProductsRepository()

	handler := &handlers.Handlers{
		GetAllProductsUsecase: &usecases.GetAllProductsImpl{
			Repo: repo,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/products", countRequests(measureRequests(handler.GetAllProducts))).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	fmt.Print("Starting server on Port:8080")
	log.Fatal(http.ListenAndServe(":8001", r))
}

func countRequests(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRequests.With(prometheus.Labels{"request": "/products"}).Inc()
		f(w, r)
	}
}

func measureRequests(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		elapsed := time.Since(start)
		requestTime.Observe(float64(elapsed))

	}
}
