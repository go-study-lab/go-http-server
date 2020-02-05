package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	//mux := http.NewServeMux()
	//mux.Handle("/", &helloHandler{})
	router := mux.NewRouter()

	router.HandleFunc("/names/{name}/countries/{country}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		name := vars["name"]
		country := vars["country"]
		fmt.Fprintf(writer, "This guy named %s, was coming from %s .", name, country)
	})

	router.Handle("/", &helloHandler{})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

