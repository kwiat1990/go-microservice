package handlers

import (
	"go-microservice/teams-api/data"
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

	// Fetch the team from the context and assign the real ID of the team
	team := r.Context().Value(KeyTeam{}).(*data.Team)
	team.ID = getTeamID(r)
	t.l.Println("[DEBUG] updating team id", team.ID)

	err := data.UpdateTeam(*team)
	if err == data.ErrTeamNotFound {
		t.l.Println("[ERROR] team not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Team not found in database"}, rw)
		return
	}

	// Write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
