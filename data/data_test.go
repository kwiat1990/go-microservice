package data

import "testing"

func TestCheckValidation(t *testing.T) {
	team := &Team{
		ShortName: "OK9",
		Name:      "Oko",
		City:      "ok",
	}
	err := team.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
