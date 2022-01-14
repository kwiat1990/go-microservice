package data

func generateNextID() int {
	t := listTeams[len(listTeams)-1]
	return t.ID + 1
}

// FindIndexByID finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByID(id int) int {
	for i, p := range listTeams {
		if p.ID == id {
			return i
		}
	}
	
	return -1
}