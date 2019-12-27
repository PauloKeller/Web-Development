package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", chien)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func chien(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
