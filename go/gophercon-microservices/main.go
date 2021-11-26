package main

import (
	"database/sql"
	"gophercon-microservices/homepage"
	"gophercon-microservices/server"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	db, err = sql.Open("mysql", "postgres://postgres:postgres@<>")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	h := homepage.NewHandlers(logger, db)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ":10443")
	logger.Println("server starting")
	err = srv.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
