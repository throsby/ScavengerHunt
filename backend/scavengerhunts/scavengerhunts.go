package scavengerhunts

import (
	"ScavengerHunt/backend/models"
	"ScavengerHunt/backend/seed_data"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var scavengerhunts []models.Hunt
var scavengerhunt_names = make(map[string]int)
var scavengerhuntIDTemp = 1

func GetScavengerHunts(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT hunt_id, title, description, username FROM Hunt INNER JOIN \"User\" ON Hunt.created_by = \"User\".user_id;")
	if err != nil {
		log.Printf("Failed to query database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hunts"})
		return
	}

	var scavengerhunts []models.Hunt
	// Iterate through results of query
	for rows.Next() {
		var hunt models.Hunt
		// Scan assigns values put into args list
		err = rows.Scan(&hunt.HuntID, &hunt.Title, &hunt.Description, &hunt.Creator)
		if err != nil {
			log.Printf("Failed to query database: %v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "No rows found during query of hunts"})
			return
		}
		scavengerhunts = append(scavengerhunts, hunt)
	}
	c.JSON(http.StatusOK, scavengerhunts)
}

func JSONGetScavengerHunts(c *gin.Context) {
	if len(scavengerhunts) == 0 {
		for _, person := range seed_data.ScavengerHuntSeed {
			person.HuntID = scavengerhuntIDTemp
			// fmt.Printf("%+v\n", person)
			createScavengerHuntByJson(c, person)
			scavengerhuntIDTemp++
		}
		return
	}
	// c.IndentedJSON(http.StatusOK, scavengerhunts)
	c.JSON(http.StatusOK, scavengerhunts)
}

func createScavengerHuntByJson(c *gin.Context, newScavengerHunt models.Hunt) {

	newScavengerHuntJSON, err := json.Marshal(newScavengerHunt)
	if err != nil {
		log.Println("New ScavengerHunt: \t", newScavengerHuntJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	scavengerhunts = append(scavengerhunts, newScavengerHunt)
	scavengerhunt_names[newScavengerHunt.Title] = newScavengerHunt.HuntID
	c.IndentedJSON(http.StatusCreated, newScavengerHunt)
}

func CreateScavengerHunt(c *gin.Context) {
	var newScavengerHunt models.Hunt
	newScavengerHunt.HuntID = scavengerhuntIDTemp
	// Attempting to add the ID field here so that there's no chance of collisions later on and can match only on username
	if err := c.BindJSON(&newScavengerHunt); err != nil {
		return
	}
	if _, exists := scavengerhunt_names[newScavengerHunt.Title]; exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "this scavengerhunt name is already being used"})
		return
	}

	scavengerhuntIDTemp++
	scavengerhunts = append(scavengerhunts, newScavengerHunt)
	scavengerhunt_names[newScavengerHunt.Title] = newScavengerHunt.HuntID
	c.IndentedJSON(http.StatusCreated, newScavengerHunt)
}

func GetScavengerHuntById(id int) (*models.Hunt, error) {
	for i, scavengerhunt := range scavengerhunts {
		if scavengerhunt.HuntID == id {
			return &scavengerhunts[i], nil
		}
	}

	return nil, errors.New("scavengerhunt not found")
}

func JSONScavengerHuntById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	scavengerhunt, err := GetScavengerHuntById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "scavenger hunt not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, scavengerhunt)
}

func ScavengerHuntById(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must use int value for hunt_id"})
		return
	}

	row := db.QueryRow("SELECT hunt_id, title, description, username FROM Hunt INNER JOIN \"User\" ON Hunt.created_by = \"User\".user_id WHERE Hunt.hunt_id = $1;", id)
	var scavengerhunt models.Hunt
	err = row.Scan(&scavengerhunt.HuntID, &scavengerhunt.Title, &scavengerhunt.Description, &scavengerhunt.Creator)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "scavenger hunt not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, scavengerhunt)
}

func AddScavengerHuntClueToHunt(c *gin.Context) {
	panic("unimplemented")
	// huntID, err := strconv.Atoi(c.Param("huntID"))
	// if err != nil {
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }
	// hunt, err := GetScavengerHuntById(huntID)
	// if err != nil {
	// 	// Handle error
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This scavenger hunt doesn't seem to exist"})
	// }
	// clueID, err := strconv.Atoi(c.Param("clueID"))
	// if err != nil {
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }
	// clue, err := scavengerhuntclues.GetScavengerHuntClueById(clueID)
	// // log.Println(clue)
	// if err != nil {
	// 	// Handle error
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This clue doesn't seem to exist"})
	// }
	//
	// clue.ScavengerHunts = append(clue.ScavengerHunts, hunt.Title)
	// hunt.ScavengerHuntClues = append(hunt.ScavengerHuntClues, *clue)
	// c.IndentedJSON(http.StatusOK, hunt)
}

func RemoveScavengerHuntClueById(c *gin.Context) {
	panic("unimplemented")
	// huntID, err := strconv.Atoi(c.Param("huntID"))
	// if err != nil {
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }
	// hunt, err := GetScavengerHuntById(huntID)
	// if err != nil {
	// 	// Handle error
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This scavenger hunt doesn't seem to exist"})
	// }
	// clueID, err := strconv.Atoi(c.Param("clueID"))
	// if err != nil {
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }
	// clue, err := scavengerhuntclues.GetScavengerHuntClueById(clueID)
	// if err != nil {
	// 	// Handle error
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This clue doesn't seem to exist"})
	// }
	//
	// // TODO Check for if the clue is in the hunt
	//
	// for _, huntToFind := range clue.ScavengerHunts {
	// 	// Confirms by name because each ScavengerHunt must have a unique name
	// 	// This is probably a bad strategy though
	// 	if hunt.Title == huntToFind {
	// 		// Remove hunt.name from []clue.scavengerhunts
	// 	}
	// }
	// for _, clueToFind := range hunt.ScavengerHuntClues {
	// 	// Confirms by ID because a clue can have several names
	// 	if clue.ClueID == clueToFind.ID {
	// 		// Remove clueToFind from []hunt.ScavengerHuntClues{}
	// 	}
	// }
	//
	// c.IndentedJSON(http.StatusOK, hunt)
}
