package users

import (
	"ScavengerHunt/backend/models"
	"ScavengerHunt/backend/seed_data"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
)

var users []models.User
var usernames = make(map[string]int)

// User Functions
var userIDTemp = 1

func GetUsers(c *gin.Context) {
	// This exists to create seed Users for the purposes of testing at the start
	// From here
	if len(users) == 0 {
		for _, person := range seed_data.UsersSeed {
			person.UserID = userIDTemp
			// fmt.Printf("%+v\n", person)
			createUserByJson(c, person)
			userIDTemp++
		}
		return
	}
	// To here

	c.IndentedJSON(http.StatusOK, users)
}

func createUserByJson(c *gin.Context, newUser models.User) {

	newUserJSON, err := json.Marshal(newUser)
	if err != nil {
		log.Println("New User: \t", newUserJSON)
		log.Println("Error marshaling JSON:", err)
		return
	}
	users = append(users, newUser)
	usernames[newUser.Username] = newUser.UserID
	c.IndentedJSON(http.StatusCreated, newUser)
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	newUser.UserID = userIDTemp
	// Attempting to add the ID field here so that there's no chance of collisions later on and can match only on username
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	if _, exists := usernames[newUser.Username]; exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "this username is already being used"})
		return
	}

	userIDTemp++
	users = append(users, newUser)
	usernames[newUser.Username] = newUser.UserID
	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUserById(id int) (*models.User, error) {
	for i, user := range users {
		if user.UserID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func UserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	user, err := GetUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
