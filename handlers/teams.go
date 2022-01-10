// Package classification Teams API
//
// Documentation for Teams API
//
// Schemes: http
// BasePath: /
// Version: 0.1.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"go-microservice/data"
	"log"
	"net/http"
)

type KeyTeam struct{}

// A list of teams
// swagger:response teamsResponse
type teamResponse struct {
	// All teams in the system
	// in: body
	Body data.Teams
}

// swagger:response noContent
type teamNoContent struct {}

// swagger:parameters deleteTeam
type teamIDParameterWrapper struct {
	// The ID of the team to delete from the system
	// in: path
	// required: true
	ID int `json:"id"` 
} 

type Teams struct {
	l *log.Logger
}

func NewTeams(l *log.Logger) *Teams {
	return &Teams{l}
}

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
