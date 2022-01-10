package handlers

import (
	"go-microservice/data"
	"net/http"
)

// swagger:route GET / teams listTeams
// Returns a list of teams
// responses:
// 	200: teamsResponse
func (t *Teams) GetTeams(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Teams")

	teams := data.GetTeams()
	err := teams.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to serve JSON", http.StatusInternalServerError)
	}
}
