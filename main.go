package main

import (
	"context"
	"example.com/http_demo/router"
	"example.com/http_demo/utils/vlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)


func main() {
	//mux := http.NewServeMux()
	//mux.Handle("/", &helloHandler{})
	muxRouter := mux.NewRouter()

	// register route handlers
	router.RegisterRoutes(muxRouter)

	// set error log writer
	errorWriter := vlog.ErrorLog.Writer()
	defer errorWriter.Close()

	server := &http.Server{
		Addr:    ":8080",
		Handler: muxRouter,
		ErrorLog: log.New(vlog.ErrorLog.Writer(), "", 0),
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

