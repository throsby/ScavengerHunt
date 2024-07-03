package scavengerhuntclues

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

var clues []models.ScavengerHuntClue
var cluenames = make(map[string]string)

// ScavengerHuntClue Functions
var clueIDTemp = 1

func GetScavengerHuntClues(c *gin.Context) {
	// This exists to create seed ScavengerHuntClues for the purposes of testing at the start
	// From here
	if len(clues) == 0 {
		for _, clue := range seed_data.ScavengerHuntCluesSeed {
			clue.ID = strconv.Itoa(clueIDTemp)
			// fmt.Printf("%+v\n", person)
			createScavengerHuntClueByJson(c, clue)
			clueIDTemp++
		}
		return
	}
	// To here

	c.IndentedJSON(http.StatusOK, clues)
}

func createScavengerHuntClueByJson(c *gin.Context, newClue models.ScavengerHuntClue) {

	newClueJSON, err := json.Marshal(newClue)
	if err != nil {
		log.Println("New ScavengerHuntClue: \t", newClueJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	clues = append(clues, newClue)
	cluenames[newClue.Name] = newClue.ID
	c.IndentedJSON(http.StatusCreated, newClue)
}

func CreateScavengerHuntClue(c *gin.Context) {
	var newClue models.ScavengerHuntClue
	newClue.ID = strconv.Itoa(clueIDTemp)
	// Attempting to add the ID field here so that there's no chance of collisions later on and can match only on cluename
	if err := c.BindJSON(&newClue); err != nil {
		return
	}
	if _, exists := cluenames[newClue.Name]; exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "this cluename is already being used"})
		return
	}

	clueIDTemp++
	clues = append(clues, newClue)
	cluenames[newClue.Name] = newClue.ID
	c.IndentedJSON(http.StatusCreated, newClue)
}

func GetScavengerHuntClueById(id string) (*models.ScavengerHuntClue, error) {
	for i, clue := range clues {
		if clue.ID == id {
			return &clues[i], nil
		}
	}

	return nil, errors.New("Clue not found")
}

func ScavengerHuntClueById(c *gin.Context) {
	id := c.Param("id")
	clue, err := GetScavengerHuntClueById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ScavengerHuntClue not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, clue)
}
