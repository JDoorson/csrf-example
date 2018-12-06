package main

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/message", ShowMessageForm)
	r.HandleFunc("/message/post", PostMessageForm).Methods("POST")

	http.ListenAndServe(":80", csrf.Protect([]byte("key"), csrf.Secure(false))(r))
}

func ShowMessageForm(w http.ResponseWriter, r *http.Request) {

}

func PostMessageForm(w http.ResponseWriter, r *http.Request) {

}