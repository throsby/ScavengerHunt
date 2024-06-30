package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
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

//
// User Functions
//

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

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

// func seedData(c *gin.Context) {
// 	for
// }

// Main

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", userById)
	router.POST("/users", createUser)

	router.GET("/teams", getTeams)
	router.GET("/teams/:id", teamById)

	router.PATCH("/teams/add/:teamID/:userID", addUserToTeamByUserID)
	router.PATCH("/teams/remove/:teamID/:userID", removeUserFromTeamByUserID)
	router.Run("localhost:8080")

}
