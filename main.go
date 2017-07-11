package main

import (
	"fmt"
	"net/http"

	"./handlers"
)

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "yoo people wassup!")
}

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/user/", handlers.UrlParams)

	http.ListenAndServe(":4080", nil)
}
