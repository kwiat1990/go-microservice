package handlers

import (
	"context"
	"fmt"
	"go-microservice/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Teams struct {
	l *log.Logger
}

func NewTeams(l *log.Logger) *Teams {
	return &Teams{l}
}

func (t *Teams) GetTeams(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Teams")

	teams := data.GetTeams()
	err := teams.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to serve JSON", http.StatusInternalServerError)
	}
}

func (t *Teams) PostTeam(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Team")

	// Cast returned value to a Team type
	team := r.Context().Value(KeyTeam{}).(data.Team)
	data.AddTeam(&team)
}

func (t *Teams) PutTeam(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID to number", http.StatusBadRequest)
		return
	}

	t.l.Printf("Handle PUT Team for ID: %d", id)

	// Cast returned value to a Team type
	team := r.Context().Value(KeyTeam{}).(data.Team)
	err = data.UpdateTeam(id, &team)
	if err == data.ErrTeamNotFound {
		http.Error(rw, "Team not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Team not found", http.StatusInternalServerError)
		return
	}
}

type KeyTeam struct{}

func (t *Teams) MiddlewareTeamValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		team := data.Team{}
		err := team.FromJSON(r.Body)
		if err != nil {
			t.l.Println("[ERROR] deserializing JSON")
			http.Error(rw, "Unable to read JSON", http.StatusBadRequest)
			return
		}

		err = team.Validate()
		if err != nil {
			t.l.Printf("[ERROR] validating team: %s\n", err)
			http.Error(
				rw,
				fmt.Sprintf("OKO! Error validating team: %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyTeam{}, team)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
