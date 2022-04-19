package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	code, err := ioutil.ReadFile("./template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t := string(code)

	http.HandleFunc("/", mainHandler(t))
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/logged-in", loggedInHandler)

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}

func mainHandler(t string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("template").Parse(t)
		if err != nil {
			fmt.Fprintf(w, "404")
		}
		tmpl.Execute(w, nil)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

}

func loggedInHandler(w http.ResponseWriter, r *http.Request) {

}
