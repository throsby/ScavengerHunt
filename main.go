package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"ScavengerHunt/backend/teams"
	"ScavengerHunt/backend/users"

	"github.com/gin-gonic/gin"
)

func addUserToTeamByUserID(c *gin.Context) {
	teamID := c.Param("teamID")
	userID := c.Param("userID")

	team, err := teams.GetTeamById(teamID)
	if err != nil {
		// Handle error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This team doesn't seem to exist"})
	}
	user, err := users.GetUserById(userID)
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

	team, teamErr := teams.GetTeamById(teamID)
	if teamErr != nil {
		// Handle error
		message := fmt.Sprintln("This team with id: ", teamID, "doesn't seem to exist")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
	}

	user, userErr := users.GetUserById(userID)
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

	router := gin.Default()
	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.UserById)
	router.POST("/users", users.CreateUser)

	router.GET("/teams", teams.GetTeams)
	router.GET("/teams/:id", teams.TeamById)

	router.PATCH("/teams/add/:teamID/:userID", addUserToTeamByUserID)
	router.PATCH("/teams/remove/:teamID/:userID", removeUserFromTeamByUserID)
	// For testing
	go hitUsersEndpoint()

	router.Run("localhost:8080")
}
