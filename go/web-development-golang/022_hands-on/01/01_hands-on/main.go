package main

import (
	"io"
	"net/http"
)

type hotdog int

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Paulo Keller")
}

func main() {
	http.HandleFunc("/", cat)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
