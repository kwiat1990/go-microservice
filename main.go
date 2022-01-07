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
	logger := log.New(os.Stdout, "go-microservice", log.LstdFlags)

	th := handlers.NewTeams(logger)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", th)

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
			logger.Fatalf("Error occured while starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Gracefully shutting down", sig)

	timeCtx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		logger.Fatal(err)
	}
	server.Shutdown(timeCtx)
}
