package main

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	EmailAddress string `json:"emailaddress"`
	Username     string `json:"username"`
}

type Team struct {
	ID                 string `json:"id"`
	TeamName           string `json:"teamname"`
	NumberOfMembers    int    `json:"numberofmembers"`
	MaxNumberOfMembers int    `json:"maxnumberofmembers"`
	TeamMembers        []User `json:"teammembers"`
}
