package handlers

import (
	"go-microservice/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Teams struct {
	l *log.Logger
}

func NewTeams(l *log.Logger) *Teams {
	return &Teams{l}
}

func (t *Teams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t.getTeams(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		t.postTeam(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		result := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(result) != 1 || len(result[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(result[0][1])
		if err != nil {
			http.Error(rw, "Invalid URL, unable to convert ID to number", http.StatusBadRequest)
			return
		} 

		t.putTeam(id, rw, r) 
		return
	}

	// Catch all unsupported methods
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (t *Teams) getTeams(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Teams")

	teams := data.GetTeams()
	err := teams.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to serve JSON", http.StatusInternalServerError)
	}
}

func (t *Teams) postTeam(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Team")

	team := &data.Team{}
	err := team.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read JSON", http.StatusBadRequest)
	}
	data.AddTeam(team)
}

func (t *Teams) putTeam(id int, rw http.ResponseWriter, r *http.Request ) {
	t.l.Println("Handle PUT Team")

	team := &data.Team{}
	err := team.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read JSON", http.StatusBadRequest)
	}

	err = data.PutTeam(id, team)
	if err == data.ErrTeamNotFound {
		http.Error(rw, "Team not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Team not found", http.StatusInternalServerError)
		return
	} 
}
