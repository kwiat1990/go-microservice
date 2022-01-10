package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-microservice/handlers"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "[go-microservice] *** ", log.LstdFlags)

	th := handlers.NewTeams(logger)

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", th.GetTeams)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", th.PutTeam)
	putRouter.Use(th.MiddlewareTeamValidation)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", th.PostTeam)
	postRouter.Use(th.MiddlewareTeamValidation)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
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
