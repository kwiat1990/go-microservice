package handlers

import (
	"go-microservice/data"
	"log"
	"net/http"
)

type Teams struct {
	l *log.Logger
}

func NewTeams(l *log.Logger) *Teams {
	return &Teams{l}
}

func (t *Teams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t.getTeams(rw, r)
		return
	}

	// Catch all unsuported methods
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (t *Teams) getTeams(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Teams")
	
	teams := data.GetTeams()
	err := teams.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to serve JSON", http.StatusInternalServerError)
	}
}
