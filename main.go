package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	scavengerhunt "ScavengerHunt/backend/scavenger_hunt"
	"ScavengerHunt/backend/teams"
	"ScavengerHunt/backend/users"

	"github.com/gin-gonic/gin"
)

func hitUsersEndpoint() {
	// Wait for 2 seconds
	time.Sleep(2 * time.Second)

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
	fmt.Println("Response:", body.String())
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

	router.GET("/scavenger_hunt", scavengerhunt.GetScavengerHunts)
	router.GET("/scavenger_hunt/:id", scavengerhunt.ScavengerHuntById)
	router.POST("/scavenger_hunt", scavengerhunt.CreateScavengerHunt)
	// For testing
	go hitUsersEndpoint()

	router.Run("localhost:8080")
}
