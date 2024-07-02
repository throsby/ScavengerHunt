package main

var usersSeed = []User{
	{FirstName: "Throsby", LastName: "Wells", EmailAddress: "twells@gmail.com", Username: "throsbyw"},
	{FirstName: "Portia", LastName: "Wells", EmailAddress: "pwells@gmail.com", Username: "portiaw"},
	{FirstName: "Tim", LastName: "Wells", EmailAddress: "t2wells@gmail.com", Username: "timw"},
	{FirstName: "Portia", LastName: "Wells", EmailAddress: "pwells@gmail.com", Username: "portiaw"},
}

var teamsSeed = []Team{
	{
		ID:                 "1",
		TeamName:           "Apricot Apricot",
		NumberOfMembers:    0,
		MaxNumberOfMembers: 4,
		TeamMembers:        []User{},
	},
}
