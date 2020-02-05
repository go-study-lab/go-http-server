package main

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	indexRouter := r.PathPrefix("/index").Subrouter()
	indexRouter.Handle("/", &HelloHandler{})

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/names/{name}/countries/{country}", ShowVisitorInfo)
}