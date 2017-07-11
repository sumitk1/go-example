package main

import (
	"fmt"
	"net/http"

	"./handlers"
)

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "yoo people wassup!")
}

//func UrlParams(w http.ResponseWriter, r *http.Request) {
//	if (r.)
//}

func main() {
	http.HandleFunc("/", echo)
	//http.HandleFunc("/user/", UrlParams)

	http.ListenAndServe(":4080", nil)
}
