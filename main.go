package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/ok", ok)
	fmt.Println("starting Server")
	http.ListenAndServe(":5000", r)
}
