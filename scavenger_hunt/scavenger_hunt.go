package scavengerhunt

import (
	"ScavengerHunt/backend/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var scavengerhunts []models.ScavengerHunt
var scavengerhunt_names = make(map[string]string)

func GetScavengerHunts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, scavengerhunts)
}

var scavengerhuntIDTemp = 1

func CreateScavengerHunt(c *gin.Context) {
	var newScavengerHunt models.ScavengerHunt
	newScavengerHunt.ID = strconv.Itoa(scavengerhuntIDTemp)
	// Attempting to add the ID field here so that there's no chance of collisions later on and can match only on username
	if err := c.BindJSON(&newScavengerHunt); err != nil {
		return
	}
	if _, exists := scavengerhunt_names[newScavengerHunt.ScavengerHuntName]; exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "this scavengerhunt name is already being used"})
		return
	}

	scavengerhuntIDTemp++
	scavengerhunts = append(scavengerhunts, newScavengerHunt)
	scavengerhunt_names[newScavengerHunt.ScavengerHuntName] = newScavengerHunt.ID
	c.IndentedJSON(http.StatusCreated, newScavengerHunt)
}

func GetScavengerHuntById(id string) (*models.ScavengerHunt, error) {
	for i, scavengerhunt := range scavengerhunts {
		if scavengerhunt.ID == id {
			return &scavengerhunts[i], nil
		}
	}

	return nil, errors.New("scavengerhunt not found")
}

func ScavengerHuntById(c *gin.Context) {
	id := c.Param("id")
	scavengerhunt, err := GetScavengerHuntById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "scavenger hunt not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, scavengerhunt)
}
