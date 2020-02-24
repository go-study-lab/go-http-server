package router

import (
	"example.com/http_demo/middleware"
	"github.com/gorilla/mux"

	"example.com/http_demo/handler"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(middleware.Logging())
	indexRouter := r.PathPrefix("/index").Subrouter()
	indexRouter.Handle("/", &handler.HelloHandler{})
	indexRouter.HandleFunc("/display_headers", handler.DisplayHeadersHandler)
	indexRouter.HandleFunc("/display_url_params", handler.DisplayUrlParamsHandler)
	indexRouter.HandleFunc("/display_form_data", handler.DisplayFormDataHandler).Methods("POST")
	indexRouter.HandleFunc("/read_cookie", handler.ReadCookieHandler)
	indexRouter.HandleFunc("/parse_json_request", handler.ParseJsonRequestHandler)

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/names/{name}/countries/{country}", handler.ShowVisitorInfo)
	userRouter.Use(middleware.Method("GET"))
}