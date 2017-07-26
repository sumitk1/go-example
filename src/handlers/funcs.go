package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
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

func postHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://www.google.com"
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
