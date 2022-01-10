package handlers

import (
	"go-microservice/data"
	"net/http"
)

func (t *Teams) PostTeam(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Team")

	// Cast returned value to a Team type
	team := r.Context().Value(KeyTeam{}).(data.Team)
	data.AddTeam(&team)
}