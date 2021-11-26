package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // sql driver
)

// DB current sql database
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://api:mysecret@postgres/auction-house?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
