package scavengerhuntclues

import (
	"ScavengerHunt/backend/models"
	"ScavengerHunt/backend/scavengerhunts"
	"ScavengerHunt/backend/seed_data"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var clues []models.Clue
var cluenames = make(map[string]int)

// Clue Functions
var clueIDTemp = 1

func GetScavengerHuntClues(c *gin.Context) {
	// This exists to create seed ScavengerHuntClues for the purposes of testing at the start
	// From here
	if len(clues) == 0 {
		for _, clue := range seed_data.ScavengerHuntCluesSeed {
			clue.ClueID = clueIDTemp
			// fmt.Printf("%+v\n", person)
			createScavengerHuntClueByJson(c, clue)
			clueIDTemp++
		}
		return
	}
	// To here
	c.IndentedJSON(http.StatusOK, clues)
}

func createScavengerHuntClueByJson(c *gin.Context, newClue models.Clue) {

	newClueJSON, err := json.Marshal(newClue)
	if err != nil {
		log.Println("New Clue: \t", newClueJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	clues = append(clues, newClue)
	cluenames[newClue.Name] = newClue.ClueID
	c.IndentedJSON(http.StatusCreated, newClue)
}

func CreateScavengerHuntClue(c *gin.Context) {
	var newClue models.Clue
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
	cluenames[newClue.Name] = newClue.ClueID
	c.IndentedJSON(http.StatusCreated, newClue)
}

func GetScavengerHuntClueById(id int) (*models.Clue, error) {
	for i, clue := range clues {
		if clue.ClueID == id {
			return &clues[i], nil
		}
	}

	return nil, errors.New("Clue not found")
}

func ScavengerHuntClueById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	clue, err := GetScavengerHuntClueById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Scavenger Hunt Clue not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, clue)
}

func WrapGetCluesByHuntID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, getCluesByHuntId(id))
}

func getCluesByHuntId(id int) /*(*[]models.Clue, error)*/ any {

	hunt, err := scavengerhunts.GetScavengerHuntById(id)
	if err != nil {
		// return nil, err

	}
	return (hunt)

	// return , nil
}

// func CluesByHuntId(c *gin.Context) {
// 	huntid, err := strconv.Atoi(c.Param("huntid"))
// 	if err != nil {
// 		c.Status(http.StatusBadRequest)
// 		return
// 	}
// 	clues := getCluesByHuntId(huntid)
// 	c.IndentedJSON(http.StatusOK, clues)
// }
