package src

import (
	"fmt"
	"io/ioutil"
	"strings"

	"html/template"
	"log"
	"net/http"

	"./handlers"
)

// OOP implementation
// only accessible inside this class - to make it global use Person
type person struct {
	fname string
	lname string
}

// Person is receiving this ServeHTTP() method
func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloooo oop func! fname = " + p.fname + " -- lname = " + p.lname))
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "yoo people wassup!")
}

type Page struct {
	Title string
	Body  []byte
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p := &Page{Title: title}
	p.Body = []byte("some good body!!")
	t, _ := template.ParseFiles("./view/template.html")
	t.Execute(w, p)
}

type IndexPageHandler struct {
}

func (this *IndexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	data, err := ioutil.ReadFile("./view/index.gohtml")
	log.Println("path = " + path)
	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	} else {
		log.Println("Err = " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", echo)
	log.Println("some log1 ... ")
	http.HandleFunc("/user/", handlers.UrlParams)
	log.Println("some log2 ... ")
	http.HandleFunc("/edit/", editHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/index", new(IndexPageHandler))

	http.HandleFunc("/post/", postHandler)

	//person_handler := &person{fname: "my_first_name", lname: "my_last_name"}
	//http.HandleFunc("/oop/", &person{fname: "my_first_name", lname: "my_last_name"})

	log.Println("some log ... ")
	http.ListenAndServe(":4080", nil)
}
