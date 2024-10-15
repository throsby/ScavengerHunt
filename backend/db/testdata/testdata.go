package main

import (
	"database/sql"
	"log"

	"ScavengerHunt/backend/models"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=throsbywells dbname=myappdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var ScavengerHuntsSeed = []models.Hunt{
		{
			HuntID:      1,
			Title:       "Gilded Age",
			MaxTeamSize: 3,
			CreatedBy:   1,
			Description: "Blurb blurb blurb",
		},
		{
			HuntID:      2,
			Title:       "Silvered Age",
			MaxTeamSize: 3,
			CreatedBy:   1,
			Description: "Blurb blurb blurb",
		},
		{
			HuntID:      3,
			Title:       "Bronzered Age",
			MaxTeamSize: 4,
			CreatedBy:   2,
			Description: "Blurb blurb blurb",
		},
		{
			HuntID:      4,
			Title:       "Gilded Age II: The Return",
			MaxTeamSize: 3,
			CreatedBy:   1,
			Description: "Blurb blurb blurb",
		},
		{
			HuntID:      5,
			Title:       "Silvered Age: Big Screen",
			MaxTeamSize: 3,
			CreatedBy:   1,
			Description: "Blurb blurb blurb",
		},
		{
			HuntID:      6,
			Title:       "Bronzered Age II: Back in the Habit",
			MaxTeamSize: 4,
			CreatedBy:   2,
			Description: "Blurb blurb blurb",
		},
		{
			HuntID:      7,
			Title:       "Gilded Cinematic Universe",
			MaxTeamSize: 3,
			CreatedBy:   3,
			Description: "Blurb blurb blurb",
		},
	}

	var TeamsSeed = []models.Team{
		{
			TeamID: 1,
			Name:   "Apricot Apricot",
			HuntID: 1,
		},
		{
			TeamID: 2,
			Name:   "Upricot Downricot",
			HuntID: 1,
		},
		{
			TeamID: 3,
			Name:   "Leftricot Rightricot",
			HuntID: 1,
		},
	}

	var ScavengerHuntCluesSeed = []models.Clue{
		{
			ClueID:         1,
			Name:           "Brooklyn Bridge",
			Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
			HuntID:         1,
			Value:          20,
			MaxSubmissions: 3,
		},
		{
			ClueID:         2,
			Name:           "Manhattan Bridge",
			Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
			HuntID:         1,
			Value:          20,
			MaxSubmissions: 3,
		},
		{
			ClueID:         3,
			Name:           "Bridge to Terabithia",
			Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
			HuntID:         1,
			Value:          20,
			MaxSubmissions: 3,
		},
		{
			ClueID:         4,
			Name:           "London Bridge",
			Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
			HuntID:         1,
			Value:          20,
			MaxSubmissions: 3,
		},
		{
			ClueID:         5,
			Name:           "II Bridge II London",
			Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
			HuntID:         2,
			Value:          20,
			MaxSubmissions: 3,
		},
		{
			ClueID:         6,
			Name:           "Sky Bridge",
			Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
			HuntID:         2,
			Value:          5,
			MaxSubmissions: 3,
		},
		{
			ClueID:         7,
			Name:           "Bifrost",
			Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
			HuntID:         2,
			Value:          15,
			MaxSubmissions: 3,
		},
	}

	var UsersSeed = []models.User{
		{Username: "Horatio", Email: "twerasdfs@gmail.com", Password: "asdflkj"},
		{Username: "Mercedes", Email: "clarlasdfs@gmail.com", Password: "asdflkj"},
		{Username: "Drakmar", Email: "t2wasdfasdfls@gmail.com", Password: "asdflkj"},
		{Username: "Groucho", Email: "garasdfble@gmail.com", Password: "asdflkj"},
	}

	// Insert users into the database
	for _, user := range UsersSeed {
		_, err := db.Exec(
			`INSERT INTO "User" (username, email, password) VALUES ($1, $2, $3)`,
			user.Username,
			user.Email,
			user.Password,
		)
		if err != nil {
			log.Printf("Error inserting user %s: %v\n", user.Username, err)
		} else {
			log.Println("Users inserted successfully.")
		}
	}

	// Insert teams into the database
	for _, hunt := range ScavengerHuntsSeed {
		_, err := db.Exec(
			`INSERT INTO Hunt (title, description, created_by) VALUES ($1, $2, $3)`,
			hunt.Title,
			hunt.Description,
			hunt.CreatedBy,
			// hunt.MaxTeamSize
		)
		if err != nil {
			log.Printf("Error inserting Scavenger Hunt %s: %v\n", hunt.Title, err)
		} else {
			log.Println("Scavenger Hunts inserted successfully.")
		}
	}

	// Insert clues into the database
	for _, clue := range ScavengerHuntCluesSeed {
		_, err := db.Exec(
			`INSERT INTO Clue (description, hunt_id, category, score) VALUES ($1, $2, $3, $4)`,
			// clue.Name,
			clue.Text,
			clue.HuntID,
			clue.Category,
			clue.Value,
		)
		if err != nil {
			log.Printf("Error inserting Scavenger Hunt Clue %s: %v\n", clue.Name, err)
		} else {
			log.Println("Hunt Clues inserted successfully.")
		}
	}

	// Insert teams into the database
	for _, team := range TeamsSeed {
		_, err := db.Exec(
			`INSERT INTO Team (name, hunt_id) VALUES ($1, $2)`,
			team.Name,
			team.HuntID,
		)
		if err != nil {
			log.Printf("Error inserting Team %s: %v\n", team.Name, err)
		} else {
			log.Println("Teams inserted successfully.")
		}
	}
}
