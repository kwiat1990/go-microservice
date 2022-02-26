package main

import (
	"context"
	"go-microservices/image-uploader/files"
	"go-microservices/image-uploader/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	basePath := "./filestore"
	l := log.New(os.Stdout, "*** Image Uploader *** ", 1)

	// create the storage class, use local storage
	// max filesize 5 MB
	storage, err := files.NewLocal(basePath, 1024*1000*5)
	if err != nil {
		l.Printf("Unable to create storage, \"error\": %s \n", err)
		os.Exit(1)
	}

	filesHandler := handlers.NewFiles(storage, l)
	middleware := handlers.GzipHandler{}
 
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))),
	)
	getRouter.Use(middleware.GzipMiddleware)
	
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", filesHandler.ServeHTTP)

	server := http.Server{
		Addr:         ":8081",
		Handler:      router,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port", server.Addr)

		err := server.ListenAndServe()
		if err != nil {
			l.Println("Unable to start server", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received.
	sig := <-c
	l.Printf("Shutting down server with \"signal\": %s\n", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
