package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"ScavengerHunt/backend/scavengerhuntclues"
	"ScavengerHunt/backend/scavengerhunts"
	"ScavengerHunt/backend/teams"
	"ScavengerHunt/backend/users"

	"github.com/gin-gonic/gin"
)

func hitEndpoints() {
	// Wait for 2 seconds
	time.Sleep(3 * time.Second)

	log.Println("\tHitting GET Users Endpoint")
	// Send the GET request
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	log.Println("Response:", body.String())

	time.Sleep(1 * time.Second)

	log.Println("\tHitting GET Teams Endpoint")
	resp, err = http.Get("http://localhost:8080/teams")
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body = new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	log.Println("Response:", body.String())

	time.Sleep(1 * time.Second)

	log.Println("\tHitting GET ScavengerHunts Endpoint")
	resp, err = http.Get("http://localhost:8080/scavengerhunts")
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body = new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	log.Println("Response:", body.String())

	time.Sleep(1 * time.Second)

	log.Println("\tHitting GET ScavengerHunt Clues Endpoint")
	resp, err = http.Get("http://localhost:8080/clues")

	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body = new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	log.Println("Response:", body.String())

	time.Sleep(1 * time.Second)

	log.Println("\tHitting Post ScavengerHuntClues Endpoint to Add clue to hunt")
	resp, err = http.Post("http://localhost:8080/scavengerhunts/1/1", "application/json", nil)
	if err != nil {
		log.Println("Error sending PATCH request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body = new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	log.Println("Response:", body.String())
}

// Main

func main() {

	router := gin.Default()
	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.UserById)
	router.POST("/users", users.CreateUser)

	router.GET("/teams", teams.GetTeams)
	router.GET("/teams/:id", teams.TeamById)

	router.PATCH("/teams/add/:teamID/:userID", teams.AddUserToTeamByUserID)
	router.PATCH("/teams/remove/:teamID/:userID", teams.RemoveUserFromTeamByUserID)

	router.GET("/scavengerhunts", scavengerhunts.GetScavengerHunts)
	router.GET("/scavengerhunts/:id", scavengerhunts.ScavengerHuntById)
	router.POST("/scavengerhunts", scavengerhunts.CreateScavengerHunt)
	router.POST("/scavengerhunts/:huntID/:clueID", scavengerhunts.AddScavengerHuntClueToHunt)

	router.GET("/clues", scavengerhuntclues.GetScavengerHuntClues)
	router.GET("/clues/:id", scavengerhuntclues.ScavengerHuntClueById)
	router.POST("/clues/mark/:id", scavengerhuntclues.MarkCorrect)

	// For testing
	go hitEndpoints()

	router.Run("localhost:8080")
}
