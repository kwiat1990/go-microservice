package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-microservice/data"
	"go-microservice/handlers"

	"github.com/go-openapi/runtime/middleware"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	newLogger := log.New(os.Stdout, "[go-microservice] *** ", log.LstdFlags)
	newValidator := data.NewValidation()

	teamsHandler := handlers.NewTeams(newLogger, newValidator)

	router := mux.NewRouter()
	router.StrictSlash(true)
	
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", teamsHandler.GetTeams)
	getRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.GetTeam)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.UpdateTeam)
	putRouter.Use(teamsHandler.MiddlewareTeamValidation)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", teamsHandler.CreateTeam)
	postRouter.Use(teamsHandler.MiddlewareTeamValidation)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", teamsHandler.DeleteTeam)

	// Serve swagger docs and load swagger spec file from disk
	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerHandler := middleware.Redoc(options, nil)

	getRouter.Handle("/docs", swaggerHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	corsHandler := gorillaHandlers.CORS(gorillaHandlers.AllowedHeaders([]string{"localhost:3000"}))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler(router),
		ErrorLog:     newLogger,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		newLogger.Println("Starting server on port 8080")
		err := server.ListenAndServe()
		if err != nil {
			newLogger.Fatalf("Error occured while starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	newLogger.Println("Gracefully shutting down", sig)

	timeCtx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		newLogger.Fatal(err)
	}
	server.Shutdown(timeCtx)
}
