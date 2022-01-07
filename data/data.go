package data

import (
	"encoding/json"
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

var listTeams = Teams{
	&Team{
		ID:        1,
		City:      "Detroit",
		Name:      "Detroit Red Wings",
		ShortName: "DET",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Team{
		ID:        2,
		City:      "Denver",
		Name:      "Colorado Avalanche",
		ShortName: "COL",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

func (t *Team) FromJSON(r io.Reader) error {
	data := json.NewDecoder(r)
	return data.Decode(t)
}

func (t *Teams) ToJSON(wr io.Writer) error {
	data := json.NewEncoder(wr)
	return data.Encode(t)
}

func GetTeams() Teams {
	return listTeams
}
