package main

import (
	"ScavengerHunt/backend/scavengerhuntclues"
	"ScavengerHunt/backend/scavengerhunts"
	"ScavengerHunt/backend/teams"
	"ScavengerHunt/backend/users"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func hitEndpoints() {
// 	// Wait for 2 seconds
// 	time.Sleep(3 * time.Second)
// 	log.Println("\tHitting GET Users Endpoint")
// 	// Send the GET request
// 	resp, err := http.Get("http://localhost:8080/users")
// 	if err != nil {
// 		log.Println("Error sending GET request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	// Read and print the response body
// 	body := new(bytes.Buffer)
// 	body.ReadFrom(resp.Body)
// 	log.Println("Response:", body.String())
// 	// time.Sleep(1 * time.Second)
// 	log.Println("\tHitting GET Teams Endpoint")
// 	resp, err = http.Get("http://localhost:8080/teams")
// 	if err != nil {
// 		log.Println("Error sending GET request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	// Read and print the response body
// 	body = new(bytes.Buffer)
// 	body.ReadFrom(resp.Body)
// 	log.Println("Response:", body.String())
// 	// time.Sleep(1 * time.Second)
// 	log.Println("\tHitting GET ScavengerHunts Endpoint")
// 	resp, err = http.Get("http://localhost:8080/scavengerhunts")
// 	if err != nil {
// 		log.Println("Error sending GET request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	// Read and print the response body
// 	body = new(bytes.Buffer)
// 	body.ReadFrom(resp.Body)
// 	log.Println("Response:", body.String())
// 	// time.Sleep(1 * time.Second)
// 	log.Println("\tHitting GET ScavengerHunt Clues Endpoint")
// 	resp, err = http.Get("http://localhost:8080/clues")
// 	if err != nil {
// 		log.Println("Error sending GET request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	// Read and print the response body
// 	body = new(bytes.Buffer)
// 	body.ReadFrom(resp.Body)
// 	log.Println("Response:", body.String())
// 	// time.Sleep(1 * time.Second)
// 	// log.Println("\tHitting Post ScavengerHuntClues Endpoint to Add clue to hunt")
// 	// resp, err = http.Post("http://localhost:8080/scavengerhunts/1/1", "application/json", nil)
// 	// if err != nil {
// 	// 	log.Println("Error sending PATCH request:", err)
// 	// 	return
// 	// }
// 	// defer resp.Body.Close()
// 	// // Read and print the response body
// 	// body = new(bytes.Buffer)
// 	// body.ReadFrom(resp.Body)
// 	// log.Println("Response:", body.String())
// }

func initDB() *sql.DB {
	// Loading the .env file so that's it's available to the os.Getenv function
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// Creating the connection string. Local_DB_USER is because I wasn't able to give permissions to another user on this machine for some reason
	connStr := "user=" + os.Getenv("LOCAL_DB_USER") + " dbname=myappdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Ping to test
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	// defer db.Close()

	return db
}

// Main

func main() {

	db := initDB()
	defer db.Close()

	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow frontend's domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
		// AllowAllOrigins:  true,
	}

	// Apply the CORS middleware to the router
	router.Use(cors.New(config))

	// router.GET("/users", users.JSONGetUsers)
	router.GET("/users", func(c *gin.Context) { users.GetUsers(c, db) })
	router.GET("/users/:id", func(c *gin.Context) { users.UserById(c, db) })
	router.POST("/users", users.CreateUser)

	router.GET("/teams", teams.GetTeams)
	router.GET("/teams/:id", teams.TeamById)

	router.PATCH("/teams/add/:teamID/:userID", teams.AddUserToTeamByUserID)
	router.PATCH("/teams/remove/:teamID/:userID", teams.RemoveUserFromTeamByUserID)

	router.GET("/scavengerhunts", func(c *gin.Context) { scavengerhunts.GetScavengerHunts(c, db) })
	router.GET("/scavengerhunts/:id", func(c *gin.Context) { scavengerhunts.ScavengerHuntById(c, db) })
	router.POST("/scavengerhunts", scavengerhunts.CreateScavengerHunt)
	router.POST("/scavengerhunts/remove/:huntID/:clueID", scavengerhunts.RemoveScavengerHuntClueById)
	router.POST("/scavengerhunts/add/:huntID/:clueID", scavengerhunts.AddScavengerHuntClueToHunt)
	router.POST("/scavengerhunts/:huntID/:clueID", scavengerhunts.AddScavengerHuntClueToHunt)

	router.GET("/clues", scavengerhuntclues.GetScavengerHuntClues)
	router.GET("/clues/:id", scavengerhuntclues.ScavengerHuntClueById)

	router.GET("test/:id", scavengerhuntclues.WrapGetCluesByHuntID)

	// router.GET("/clues/:huntid", scavengerhuntclues.CluesByHuntId)

	// For testing
	// go hitEndpoints()

	log.Fatal(router.Run("localhost:8080"))
	// log.Fatal(router.Run("0.0.0.0:8080"))
}
