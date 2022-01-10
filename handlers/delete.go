package handlers

import (
	"go-microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /{id} teams deleteTeam
// Returns void 
// responses:
// 	201: noContent
func (t *Teams) DeleteTeam(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	t.l.Println("Handle DELETE Team", id)
	err := data.DeleteTeam(id)

	if err == data.ErrTeamNotFound {
		http.Error(rw, "Team not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Team not found", http.StatusInternalServerError)
		return
	}
}
