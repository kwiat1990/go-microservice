package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-microservice/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "[go-microservice] *** ", log.LstdFlags)

	teamsHandler := handlers.NewTeams(logger)

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", teamsHandler.GetTeams)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.PutTeam)
	putRouter.Use(teamsHandler.MiddlewareTeamValidation)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", teamsHandler.PostTeam)
	postRouter.Use(teamsHandler.MiddlewareTeamValidation)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.DeleteTeam)

	// Serve swagger docs and load swagger spec file from disk
	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerHandler := middleware.Redoc(options, nil)
	
	getRouter.Handle("/docs", swaggerHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port 8080")
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatalf("Error occured while starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
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
