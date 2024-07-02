package models

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

type ScavengerHunt struct {
	ID                 string              `json:"id"`
	NumberOfTeams      int                 `json:"numberofteams"`
	MaxNumberOfTeams   int                 `json:"maxnumberofteams"`
	Teams              []Team              `json:"teamms"`
	ScavengerHuntClues []ScavengerHuntClue `json:"scavengerhuntclues"`
}

type ScavengerHuntClue struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Text          string        `json:"text"`
	PointValue    string        `json:"pointvalue"`
	ScavengerHunt ScavengerHunt `json:"scavengerhunts"`
}
