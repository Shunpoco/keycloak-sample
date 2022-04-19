package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Env struct {
	Login        string
	Token        string
	UserInfo     string
	ClientId     string
	ClientSecret string
}

func main() {
	code, err := ioutil.ReadFile("./template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t := string(code)
	base := "http://localhost:8080/realms/test/protocol/openid-connect"
	env := Env{
		Login:        fmt.Sprintf("%s/auth", base),
		Token:        fmt.Sprintf("%s/token", base),
		UserInfo:     fmt.Sprintf("%s/userinfo", base),
		ClientId:     "",
		ClientSecret: "",
	}

	http.HandleFunc("/", mainHandler(t, &env))
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/logged-in", loggedInHandler)

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}

func mainHandler(t string, env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("template").Parse(t)
		if err != nil {
			fmt.Fprintf(w, "404")
		}
		tmpl.Execute(w, env)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

}

func loggedInHandler(w http.ResponseWriter, r *http.Request) {

}
