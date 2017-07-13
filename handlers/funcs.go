package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func UrlParams(w http.ResponseWriter, r *http.Request) {
	// spit out the path
	fmt.Fprintln(w, r.URL.Path[:])
}

type Page struct {
	Title string
	Body  []byte
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p := &Page{Title: title}
	p.Body = []byte("some good body!!")
	t, _ := template.ParseFiles("../view/template.html")
	t.Execute(w, p)
}
