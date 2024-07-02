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

var teamnames = make(map[string]string)

//
// Team Functions
//

var teamIDTemp = 1

func GetTeams(c *gin.Context) {
	if len(teams) == 0 {
		for _, team := range seed_data.TeamsSeed {
			team.ID = strconv.Itoa(teamIDTemp)
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
	teamnames[newTeam.TeamName] = newTeam.ID
	c.IndentedJSON(http.StatusCreated, newTeam)
}

func GetTeamById(id string) (*models.Team, error) {
	for i, team := range teams {
		if team.ID == id {
			return &teams[i], nil
		}
	}

	return nil, errors.New("team not found")
}

func TeamById(c *gin.Context) {
	id := c.Param("id")
	team, err := GetTeamById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Team not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, team)
}
