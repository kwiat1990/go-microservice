package data

import (
	"fmt"
	"time"
)

// Team defines a structure for an API team
// swagger:model
type Team struct {
	ID int `json:"id"`
	// Name of the team
	//
	// required: true
	// pattern: /[a-zA-Z]+/
	Name string `json:"name" validate:"required"`
	// Shorthand name of the team
	//
	// required: true
	// pattern: /[A-Z]{3}/
	// minimum length: 3
	// maximum length: 3
	ShortName string `json:"shortName" validate:"required,uppercase,alpha"`
	// City of the team
	//
	// required: true
	City      string `json:"city" validate:"required"`
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

func GetTeamByID(id int) (*Team, error) {
	i := findIndexByID(id)
	if id == -1 {
		return nil, ErrTeamNotFound
	}

	return listTeams[i], nil
}

func AddTeam(t Team) {
	t.ID = generateNextID()
	listTeams = append(listTeams, &t)
}

func UpdateTeam(t Team) error {
	i := findIndexByID(t.ID)
	if i == -1 {
		return ErrTeamNotFound
	}

	listTeams[i] = &t

	return nil
}

func DeleteTeam(id int) error {
	i := findIndexByID(id)
	if i == -1 {
		return ErrTeamNotFound
	}

	listTeams = append(listTeams[:i], listTeams[i+1])

	return nil
}

