package scavengerhunts

import (
	"ScavengerHunt/backend/models"
	"ScavengerHunt/backend/scavengerhuntclues"
	"ScavengerHunt/backend/seed_data"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var scavengerhunts []models.ScavengerHunt
var scavengerhunt_names = make(map[string]string)
var scavengerhuntIDTemp = 1

func GetScavengerHunts(c *gin.Context) {
	if len(scavengerhunts) == 0 {
		for _, person := range seed_data.ScavengerHuntSeed {
			person.ID = strconv.Itoa(scavengerhuntIDTemp)
			// fmt.Printf("%+v\n", person)
			createScavengerHuntByJson(c, person)
			scavengerhuntIDTemp++
		}
		return
	}
	c.IndentedJSON(http.StatusOK, scavengerhunts)
}

func createScavengerHuntByJson(c *gin.Context, newScavengerHunt models.ScavengerHunt) {

	newScavengerHuntJSON, err := json.Marshal(newScavengerHunt)
	if err != nil {
		log.Println("New ScavengerHunt: \t", newScavengerHuntJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	scavengerhunts = append(scavengerhunts, newScavengerHunt)
	scavengerhunt_names[newScavengerHunt.ScavengerHuntName] = newScavengerHunt.ID
	c.IndentedJSON(http.StatusCreated, newScavengerHunt)
}

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

func AddScavengerHuntClueToHunt(c *gin.Context) {
	huntID := c.Param("huntID")
	clueID := c.Param("clueID")

	hunt, err := GetScavengerHuntById(huntID)
	if err != nil {
		// Handle error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This scavenger hunt doesn't seem to exist"})
	}
	clue, err := scavengerhuntclues.GetScavengerHuntClueById(clueID)
	log.Println(clue)
	if err != nil {
		// Handle error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This clue doesn't seem to exist"})
	}

	hunt.ScavengerHuntClues = append(hunt.ScavengerHuntClues, *clue)
	clue.ScavengerHunts = append(clue.ScavengerHunts, *hunt)
	c.IndentedJSON(http.StatusOK, hunt)
}
