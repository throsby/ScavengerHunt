package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var users []User
var teams []Team
var usernames = make(map[string]string)
var teamnames = make(map[string]string)

func addUserToTeamByUserID(c *gin.Context) {
	teamID := c.Param("teamID")
	userID := c.Param("userID")

	team, err := getTeamById(teamID)
	if err != nil {
		// Handle error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This team doesn't seem to exist"})
	}
	user, err := getUserById(userID)
	if err != nil {
		// Handle error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This user doesn't seem to exist"})
	}
	fmt.Println(team.TeamMembers)
	if len(team.TeamMembers) > 0 {
		for _, member := range team.TeamMembers {
			if member.ID == user.ID {
				c.IndentedJSON(http.StatusConflict, gin.H{"message": "This user is already a member of this team"})
				return
			}
		}
	}
	if len(team.TeamMembers) > team.MaxNumberOfMembers {
		message := fmt.Sprint("The number of team members is already at its maximum of ", team.MaxNumberOfMembers)
		c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
		return
	}
	team.TeamMembers = append(team.TeamMembers, *user)

	c.IndentedJSON(http.StatusOK, team)
}

func removeUserFromTeamByUserID(c *gin.Context) {
	teamID := c.Param("teamID")
	userID := c.Param("userID")

	team, teamErr := getTeamById(teamID)
	if teamErr != nil {
		// Handle error
		message := fmt.Sprintln("This team with id: ", teamID, "doesn't seem to exist")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
	}

	user, userErr := getUserById(userID)
	if userErr != nil {
		// Handle error
		message := fmt.Sprintln("This team with id: ", userID, "doesn't seem to exist")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
	}

	var userPosition int
	// Check if there are more than 0 teamMembers
	if len(team.TeamMembers) > 0 {
		for i, member := range team.TeamMembers {
			if member.ID == user.ID {
				userPosition = i
			}
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Team has no members to be removed"})
		return
	}

	team.TeamMembers = append(team.TeamMembers[:userPosition], team.TeamMembers[userPosition+1:]...)

	c.IndentedJSON(http.StatusOK, team)
}

//
// Team Functions
//

var teamIDTemp = 1

func getTeams(c *gin.Context) {
	if len(teams) == 0 {
		for _, team := range teamsSeed {
			team.ID = strconv.Itoa(teamIDTemp)
			// fmt.Printf("%+v\n", person)
			createTeamByJson(c, team)
			teamIDTemp++
		}
		return
	}

	c.IndentedJSON(http.StatusOK, teams)
}

func createTeamByJson(c *gin.Context, newTeam Team) {

	newTeamJSON, err := json.Marshal(newTeam)
	if err != nil {
		log.Println("New Team: \t", newTeamJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	teams = append(teams, newTeam)
	usernames[newTeam.TeamName] = newTeam.ID
	c.IndentedJSON(http.StatusCreated, newTeam)
}

func getTeamById(id string) (*Team, error) {
	for i, team := range teams {
		if team.ID == id {
			return &teams[i], nil
		}
	}

	return nil, errors.New("team not found")
}

func teamById(c *gin.Context) {
	id := c.Param("id")
	team, err := getTeamById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Team not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, team)
}

// User Functions
var userIDTemp = 1

func getUsers(c *gin.Context) {
	// This exists to create seed Users for the purposes of testing at the start
	// From here
	if len(users) == 0 {
		for _, person := range usersSeed {
			person.ID = strconv.Itoa(userIDTemp)
			// fmt.Printf("%+v\n", person)
			createUserByJson(c, person)
			userIDTemp++
		}
		return
	}
	// To here

	c.IndentedJSON(http.StatusOK, users)
}

func createUserByJson(c *gin.Context, newUser User) {

	newUserJSON, err := json.Marshal(newUser)
	if err != nil {
		log.Println("New User: \t", newUserJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	users = append(users, newUser)
	usernames[newUser.Username] = newUser.ID
	c.IndentedJSON(http.StatusCreated, newUser)
}

func createUser(c *gin.Context) {
	var newUser User
	newUser.ID = strconv.Itoa(userIDTemp)
	// Attempting to add the ID field here so that there's no chance of collisions later on and can match only on username
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	if _, exists := usernames[newUser.Username]; exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "this username already exists"})
		return
	}

	userIDTemp++
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserById(id string) (*User, error) {
	for i, user := range users {
		if user.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func userById(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func hitUsersEndpoint() {
	// Wait for 2 seconds
	time.Sleep(2 * time.Second)

	// Send the GET request
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	fmt.Println("Response:", body.String())
}

// Main

func main() {
	// fmt.Println(users, teams)
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", userById)
	router.POST("/users", createUser)

	router.GET("/teams", getTeams)
	router.GET("/teams/:id", teamById)

	router.PATCH("/teams/add/:teamID/:userID", addUserToTeamByUserID)
	router.PATCH("/teams/remove/:teamID/:userID", removeUserFromTeamByUserID)
	// For testing
	go hitUsersEndpoint()

	router.Run("localhost:8080")
}
