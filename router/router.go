package router

import (
	"github.com/gorilla/mux"

	"example.com/http_demo/handler"
)

func RegisterRoutes(r *mux.Router) {
	indexRouter := r.PathPrefix("/index").Subrouter()
	indexRouter.Handle("/", &handler.HelloHandler{})

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/names/{name}/countries/{country}", handler.ShowVisitorInfo)
}