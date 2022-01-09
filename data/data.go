package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	City      string `json:"city"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Teams []*Team

var ErrTeamNotFound = fmt.Errorf("Team not found")

var listTeams = Teams{
	&Team{
		ID:        0,
		City:      "Detroit",
		Name:      "Detroit Red Wings",
		ShortName: "DET",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Team{
		ID:        1,
		City:      "Denver",
		Name:      "Colorado Avalanche",
		ShortName: "COL",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

func GetTeams() Teams {
	return listTeams
}

func AddTeam(t *Team) {
	t.ID = generateNextID()
	listTeams = append(listTeams, t)
}

func UpdateTeam(id int, team *Team) error {
	_, i, err := findTeam(id)
	if err != nil {
		fmt.Printf("unable to find a team with the ID %d: %s", i, err)
	}
	team.ID = id
	listTeams[i] = team
	return err
}

func (t *Team) FromJSON(r io.Reader) error {
	// While working with io.Reader/Writer it's better to use json.NewEncoder/NewDecoder 
	// instead of json.Marshal/Unmarshal as it's slightly more performant.
	data := json.NewDecoder(r)
	return data.Decode(t)
}

func (t *Teams) ToJSON(wr io.Writer) error {
	data := json.NewEncoder(wr)
	return data.Encode(t)
}

func generateNextID() int {
	t := listTeams[len(listTeams)-1]
	return t.ID + 1
}

func findTeam(id int) (*Team, int, error) {
	for i, v := range listTeams {
		if v.ID == id {
			return v, i, nil
		}
	}
	return nil, -1, ErrTeamNotFound
}
