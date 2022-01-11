package handlers

import (
	"go-microservice/data"
	"net/http"
)

// swagger:route PUT /teams teams updateTeam
// Update a team details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update teams
func (t *Teams) UpdateTeam(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	// Fetch the team from the context
	prod := r.Context().Value(KeyTeam{}).(data.Team)
	t.l.Println("[DEBUG] updating team id", prod.ID)

	err := data.UpdateTeam(prod)
	if err == data.ErrTeamNotFound {
		t.l.Println("[ERROR] team not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Team not found in database"}, rw)
		return
	}

	// Write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
