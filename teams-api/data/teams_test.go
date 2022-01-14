package data

import "testing"

func TestCheckValidation(t *testing.T) {
	team := &Team{
		ShortName: "DET",
		Name:      "Detroit Red Wings",
		City:      "Detroit",
	}

	v := NewValidation()
	err := v.Validate(team)
	if err != nil {
		t.Fatal(err)
	}
}
