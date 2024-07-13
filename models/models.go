package models

type User struct {
	ID           string   `json:"id"`
	FirstName    string   `json:"firstname"`
	LastName     string   `json:"lastname"`
	EmailAddress string   `json:"emailaddress"`
	Username     string   `json:"username"`
	HuntHost     []string `json:"hunthost"`
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
	ScavengerHuntName  string              `json:"scavengerhuntname"`
	NumberOfTeams      int                 `json:"numberofteams"`
	MaxNumberOfTeams   int                 `json:"maxnumberofteams"`
	Teams              []Team              `json:"teams"`
	ScavengerHuntClues []ScavengerHuntClue `json:"scavengerhuntclues"`
}

type ScavengerHuntClue struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Text             string   `json:"text"`
	PointValue       int      `json:"pointvalue"`
	ScavengerHunts   []string `json:"scavengerhunts"`
	ConfirmedCorrect bool     `json:"confirmedcorrect"`
	ClueCode         string   `json:"cluecode"`
}
