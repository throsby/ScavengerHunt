package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	scavengerhunt "ScavengerHunt/backend/scavenger_hunt"
	"ScavengerHunt/backend/teams"
	"ScavengerHunt/backend/users"

	"github.com/gin-gonic/gin"
)

func hitEndpoints() {
	// Wait for 2 seconds
	time.Sleep(1 * time.Second)

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
	log.Println("Response:", body.String(), "\n")

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
	log.Println("Response:", body.String(), "\n")

	time.Sleep(1 * time.Second)

	log.Println("\tHitting GET Teams Endpoint")
	resp, err = http.Get("http://localhost:8080/scavenger_hunts")
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body = new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	log.Println("Response:", body.String(), "\n")
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

	router.GET("/scavenger_hunts", scavengerhunt.GetScavengerHunts)
	router.GET("/scavenger_hunts/:id", scavengerhunt.ScavengerHuntById)
	router.POST("/scavenger_hunts", scavengerhunt.CreateScavengerHunt)
	// For testing
	go hitEndpoints()

	router.Run("localhost:8080")
}
