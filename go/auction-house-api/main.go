package main

import (
	"auction-house/src/auctions"
	"auction-house/src/items"
	"auction-house/src/logger"
	"auction-house/src/server"
	"auction-house/src/users"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

var err error

// @contact.name Paulo Keller

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	r.HandleFunc("/users", users.FetchUsers).Methods("GET")
	r.HandleFunc("/users", users.CreateUser).Methods("POST")
	r.HandleFunc("/users/{user_id}", users.FetchUserByID).Methods("GET")
	r.HandleFunc("/items", items.FetchItems).Methods("GET")
	r.HandleFunc("/items", items.CreateItem).Methods("POST")
	r.HandleFunc("/items/{item_id}", items.DeleteItem).Methods("DELETE")
	r.HandleFunc("/auctions", auctions.FetchAuctions).Methods("GET")
	r.HandleFunc("/auctions", auctions.CreateAuction).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	srv := server.NewServer(":8080", handler)

	err = srv.ListenAndServe()

	if err != nil {
		logger.MongoLogger.Fatalf("server failed to start: %v", err)
	}

	if err != nil {
		logger.MongoLogger.Fatalf("server failed to start: %v", err)
	}
}
