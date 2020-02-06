package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ShowVisitorInfo(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	country := vars["country"]
	fmt.Fprintf(writer, "This guy named %s, was coming from %s .", name, country)
}