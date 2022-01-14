package handlers

import (
	"go-microservice/teams-api/data"
	"net/http"
)

// swagger:route GET / teams getTeams
// Return a list of teams from the database
// responses:
//	200: teamsResponse

// GetTeams handles GET requests and returns all current teams
func (t *Teams) GetTeams(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("[DEBUG] get all teams")
	rw.Header().Add("Content-Type", "application/json")

	teams := data.GetTeams()
	err := data.ToJSON(teams, rw)
	if err != nil {
		// We should never be here but log the error just incase
		t.l.Println("[ERROR] serializing teams", err)
	}
}

// swagger:route GET /{id} teams getTeam
// Return a single team
// responses:
//	200: teamResponse
//	404: errorResponse

// GetTeam handles GET requests and returns a single team
func (t *Teams) GetTeam(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getTeamID(r)

	t.l.Println("[DEBUG] get record id", id)

	team, err := data.GetTeamByID(id)

	switch err {
	case nil:

	case data.ErrTeamNotFound:
		t.l.Println("[ERROR] fetching team", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return

	default:
		t.l.Println("[ERROR] fetching team", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(team, rw)
	if err != nil {
		// We should never be here but log the error just incase
		t.l.Println("[ERROR] serializing team", err)
	}
}
