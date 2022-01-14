package handlers

import (
	"go-microservice/teams-api/data"
	"net/http"
)

// swagger:route DELETE /teams/{id} teams deleteTeam
// Update a teams details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (t *Teams) DeleteTeam(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getTeamID(r)

	t.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteTeam(id)
	if err == data.ErrTeamNotFound {
		t.l.Println("[ERROR] deleting team id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		t.l.Println("[ERROR] deleting team", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
