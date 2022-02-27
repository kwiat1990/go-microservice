package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-microservice/data"
	"go-microservice/handlers"

	"github.com/go-openapi/runtime/middleware"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "[go-microservice] *** ", log.LstdFlags)
	validator := data.NewValidation()
	teamsHandler := handlers.NewTeams(logger, validator)

	router := mux.NewRouter().StrictSlash(true)
	
	getRouter := router.PathPrefix("/api").Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("", teamsHandler.GetTeams)
	getRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.GetTeam)

	putRouter := router.PathPrefix("/api").Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.UpdateTeam)
	putRouter.Use(teamsHandler.MiddlewareTeamValidation)

	postRouter := router.PathPrefix("/api").Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", teamsHandler.CreateTeam)
	postRouter.Use(teamsHandler.MiddlewareTeamValidation)

	deleteRouter := router.PathPrefix("/api").Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.DeleteTeam)

	// Serve swagger docs and load swagger spec file from disk
	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerHandler := middleware.Redoc(options, nil)

	router.Handle("/docs", swaggerHandler)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	corsHandler := gorillaHandlers.CORS(gorillaHandlers.AllowedHeaders([]string{"*"}))
        
        serverPort := os.Getenv("PORT")
        if len(port) == 0 {
	    port = "8080"
        }

	server := &http.Server{
		Addr:         ":" + serverPort,
		Handler:      corsHandler(router),
		ErrorLog:     logger,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
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
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	logger.Println("Gracefully shutting down", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown Failed", "error", err)
	} else {
		logger.Panicln("Server Shutdown gracefully")
	}
}
