package teams

import (
	"ScavengerHunt/backend/models"
	"ScavengerHunt/backend/seed_data"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var teams []models.Team

var teamnames = make(map[string]int)

//
// Team Functions
//

var teamIDTemp = 1

func GetTeams(c *gin.Context) {
	if len(teams) == 0 {
		for _, team := range seed_data.TeamsSeed {
			team.TeamID = teamIDTemp
			// fmt.Printf("%+v\n", person)
			createTeamByJson(c, team)
			teamIDTemp++
		}
		return
	}

	c.IndentedJSON(http.StatusOK, teams)
}

func createTeamByJson(c *gin.Context, newTeam models.Team) {

	newTeamJSON, err := json.Marshal(newTeam)
	if err != nil {
		log.Println("New Team: \t", newTeamJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	teams = append(teams, newTeam)
	teamnames[newTeam.Name] = newTeam.TeamID
	c.IndentedJSON(http.StatusCreated, newTeam)
}

func GetTeamById(id int) (*models.Team, error) {
	for i, team := range teams {
		if team.TeamID == id {
			return &teams[i], nil
		}
	}

	return nil, errors.New("team not found")
}

func TeamById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	team, err := GetTeamById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Team not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, team)
}

func AddUserToTeamByUserID(c *gin.Context) {
	panic("unimplemented")
	// teamID, err := strconv.Atoi(c.Param("teamID"))
	// if err != nil {
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }
	// team, err := GetTeamById(teamID)
	// if err != nil {
	// 	// Handle error
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This team doesn't seem to exist"})
	// }
	// userID, err := strconv.Atoi(c.Param("userID"))
	// if err != nil {
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }
	// user, err := users.GetUserById(userID)
	// if err != nil {
	// 	// Handle error
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This user doesn't seem to exist"})
	// }

	// if len(team.TeamMembers) > 0 {
	// 	for _, member := range team.TeamMembers {
	// 		if member.userID == user.UserID {
	// 			c.IndentedJSON(http.StatusConflict, gin.H{"message": "This user is already a member of this team"})
	// 			return
	// 		}
	// 	}
	// }
	// if len(team.TeamMembers) > team.MaxNumberOfMembers {
	// 	message := fmt.Sprint("The number of team members is already at its maximum of ", team.MaxNumberOfMembers)
	// 	c.IndentedJSON(http.StatusConflict, gin.H{"message": message})
	// 	return
	// }
	// team.TeamMembers = append(team.TeamMembers, *user)

	// c.IndentedJSON(http.StatusOK, team)
}

func RemoveUserFromTeamByUserID(c *gin.Context) {
	panic("unimplemented")
	// teamID := c.Param("teamID")
	// userID := c.Param("userID")

	// team, teamErr := GetTeamById(teamID)
	// if teamErr != nil {
	// 	// Handle error
	// 	message := fmt.Sprintln("This team with id: ", teamID, "doesn't seem to exist")
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
	// }

	// user, userErr := users.GetUserById(userID)
	// if userErr != nil {
	// 	// Handle error
	// 	message := fmt.Sprintln("This team with id: ", userID, "doesn't seem to exist")
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
	// }

	// var userPosition int
	// // Check if there are more than 0 teamMembers
	// if len(team.TeamMembers) > 0 {
	// 	for i, member := range team.TeamMembers {
	// 		if member.ID == user.ID {
	// 			userPosition = i
	// 		}
	// 	}
	// } else {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Team has no members to be removed"})
	// 	return
	// }

	// team.TeamMembers = append(team.TeamMembers[:userPosition], team.TeamMembers[userPosition+1:]...)
	//
	// c.IndentedJSON(http.StatusOK, team)
}
