package handlers

import (
	"fmt"
	"net/http"
)

func UrlParams(w http.ResponseWriter, r *http.Request) {
	// spit out the path
	fmt.Fprintln(w, r.URL.Path[:])
}
