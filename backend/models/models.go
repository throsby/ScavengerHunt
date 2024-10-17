package models

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Hunt struct {
	HuntID      int    `json:"hunt_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
	Creator     string `json:"creator"`
	MaxTeamSize uint   `json:"max_team_size"`
}

type Host struct {
	HostID int `json:"host_id"`
	UserID int `json:"user_id"`
	HuntID int `json:"hunt_id"`
}

type Clue struct {
	ClueID         int    `json:"clue_id"`
	Name           string `json:"name"`
	Text           string `json:"text"`
	HuntID         int    `json:"hunt_id"`
	Category       string `json:"category"`
	Value          int    `json:"value"`
	MaxSubmissions int    `json:"max_submissions"`
}

type Team struct {
	TeamID int    `json:"team_id"`
	Name   string `json:"name"`
	HuntID int    `json:"hunt_id"`
}

type Submission struct {
	SubmissionID int    `json:"submission_id"`
	TeamID       int    `json:"team_id"`
	ClueID       int    `json:"clue_id"`
	Answer       string `json:"answer"`
	Score        int    `json:"score"`
}
