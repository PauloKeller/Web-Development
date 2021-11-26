package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat")
}

func createTemplate(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "template.gohtml", "Paulo Keller")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(cat))
	http.HandleFunc("/dog", http.HandlerFunc(dog))
	http.HandleFunc("/me/", http.HandlerFunc(createTemplate))

	http.ListenAndServe(":8080", nil)
}
