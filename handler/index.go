package handler

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (*HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World1")
}