package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

var (
	page = Page{Title: "Main page", Body: "Hi!"}
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("templates/*"))

	templates.ExecuteTemplate(w, "index", page)
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public/", files))
	mux.HandleFunc("/", indexHandler)
	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: mux,
	}
	server.ListenAndServe()
}
