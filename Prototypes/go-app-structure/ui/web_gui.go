package ui

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("config.html")
	p := Page{Title: "", Body: ""}
	t.Execute(w, &p)
}

func StartUIServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/config", configHandler)
	fmt.Println("Web GUI server starting on :5555")
	err := http.ListenAndServe(":5555", nil)
	fmt.Println(err);
}
