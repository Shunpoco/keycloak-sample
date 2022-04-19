package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/logged-in", loggedInHandler)

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

}

func loggedInHandler(w http.ResponseWriter, r *http.Request) {

}
