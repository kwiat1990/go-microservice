package handlers

import (
	"go-microservice/teams-api/data"
	"net/http"
)

// swagger:route POST /teams teams createTeam
// Create a new team
//
// responses:
//	200: teamResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new team
func (t *Teams) CreateTeam(rw http.ResponseWriter, r *http.Request) {
	// Fetch the team from the context
	team := r.Context().Value(KeyTeam{}).(*data.Team)
	t.l.Printf("[DEBUG] Inserting team: %#v\n", team)
	data.AddTeam(*team)
}
