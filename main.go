package main

import (
	"context"
	"go-microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "go-micro-service", log.LstdFlags)

	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodbye(logger)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", hh)
	serverMux.Handle("/goodbye", gh)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Gracefully shutting down", sig)

	timeCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeCtx)
}
