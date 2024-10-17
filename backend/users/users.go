package users

import (
	"ScavengerHunt/backend/models"
	"ScavengerHunt/backend/seed_data"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var users []models.User
var usernames = make(map[string]int)

// User Functions
var userIDTemp = 1

func JSONGetUsers(c *gin.Context) {
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

func GetUsers(c *gin.Context, db *sql.DB) {
	// This exists to create seed Users for the purposes of testing at the start
	// From here
	rows, err := db.Query("SELECT user_id, username, email FROM \"User\"")
	if err != nil {
		log.Fatalf("Failed to query database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	var users []models.User
	// Iterate through result of query
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.UserID, &user.Username, &user.Email)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "No rows found during query of users"})
		}
		// print(user.Username)
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
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

// TODO
func CreateUser(c *gin.Context) {
	// TODO
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

func JSONCreateUser(c *gin.Context) {
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

func JSONGetUserById(id int) (*models.User, error) {
	for i, user := range users {
		if user.UserID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

// TODO
func GetUserById(id int) (*models.User, error) {

	for i, user := range users {
		if user.UserID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func UserById(c *gin.Context, db *sql.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error: Must use int value when calling UserById because user_id field is a non-float int")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must use int value for user_id"})
		return
	}

	// Format query string
	row := db.QueryRow("SELECT user_id, username, email FROM \"User\" WHERE user_id = $1;", id)
	var user models.User
	// Assign to user
	err = row.Scan(&user.UserID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		log.Printf("Error: No row was found, this is possibly because UserById requested a user_id that doesn't exist. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func JSONUserById(c *gin.Context) {
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
