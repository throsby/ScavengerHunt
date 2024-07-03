package seed_data

import (
	"ScavengerHunt/backend/models"
)

var UsersSeed = []models.User{
	{FirstName: "Throsby", LastName: "Wells", EmailAddress: "twells@gmail.com", Username: "throsbyw"},
	{FirstName: "Portia", LastName: "Wells", EmailAddress: "pwells@gmail.com", Username: "portiaw"},
	{FirstName: "Tim", LastName: "Wells", EmailAddress: "t2wells@gmail.com", Username: "timw"},
	{FirstName: "Portia", LastName: "Wells", EmailAddress: "pwells@gmail.com", Username: "portiaw2"},
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
		ScavengerHuntName: "Gilded Age,",
		NumberOfTeams:     0,
		MaxNumberOfTeams:  15,
		Teams:             []models.Team{},
	},
}
