package handlers

import (
	"fmt"
	"go-microservice/teams-api/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyTeam struct{}

type Teams struct {
	l *log.Logger
	v *data.Validation
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// ErrInvalidTeamPath is an error message when the team path is not valid
var ErrInvalidTeamPath = fmt.Errorf("invalid Path, path should be /api/[id]")

func NewTeams(l *log.Logger, v *data.Validation) *Teams {
	return &Teams{l, v}
}

// getTeamID returns the team ID from the URL
// Panics if cannot convert the ID into an integer
// this should never happen as the router ensures that
// this is a valid number
func getTeamID(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Should never happen
		panic(err)
	}

	return id
}
