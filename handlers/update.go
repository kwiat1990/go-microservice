package handlers

import (
	"go-microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
