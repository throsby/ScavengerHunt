package seed_data

import (
	"ScavengerHunt/backend/models"
)

var UsersSeed = []models.User{
	{FirstName: "Horatio", LastName: "Puits", EmailAddress: "twerasdfs@gmail.com", Username: "Horatio"},
	{FirstName: "Mercedes", LastName: "Puits", EmailAddress: "clarlasdfs@gmail.com", Username: "Mercedes"},
	{FirstName: "Drakmar", LastName: "Puits", EmailAddress: "t2wasdfasdfls@gmail.com", Username: "Drakmar"},
	{FirstName: "Groucho", LastName: "Puits", EmailAddress: "garasdfble@gmail.com", Username: "Groucho"},
}

var TeamsSeed = []models.Team{
	{
		ID:                 "1",
		TeamName:           "Apricot Apricot",
		NumberOfMembers:    0,
		MaxNumberOfMembers: 4,
		TeamMembers:        []models.User{},
	},
}

var ScavengerHuntSeed = []models.ScavengerHunt{
	{
		ScavengerHuntName:  "Gilded Age",
		NumberOfTeams:      0,
		MaxNumberOfTeams:   15,
		Teams:              []models.Team{},
		ScavengerHuntClues: []models.ScavengerHuntClue{},
	},
}

var ScavengerHuntCluesSeed = []models.ScavengerHuntClue{
	{
		Name:             "Brooklyn Bridge",
		Text:             "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		PointValue:       25,
		ScavengerHunts:   []string{},
		ConfirmedCorrect: false,
	},
}
