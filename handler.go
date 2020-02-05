package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HelloHandler struct{}

func (*HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func ShowVisitorInfo(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	country := vars["country"]
	fmt.Fprintf(writer, "This guy named %s, was coming from %s .", name, country)
}