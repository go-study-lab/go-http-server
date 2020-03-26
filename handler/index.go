package handler

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (*HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//ints := []int{0, 1, 2}
	//fmt.Fprintf(w, "%v", ints[0:5])
	fmt.Fprintf(w, "Hello World1")
}