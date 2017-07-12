package main

import (
	"fmt"
	"net/http"

	"./handlers"
)

// OOP implementation
type person struct {
	fname string
	lname string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloooo oop func! fname = " + p.fname + " -- lname = " + p.lname))
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "yoo people wassup!")
}

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/user/", handlers.UrlParams)

	person_handler := &person{fname: "my_first_name", lname: "my_last_name"}
	//http.HandleFunc("/oop/", person_func)

	http.ListenAndServe(":4080", person_handler)
}
