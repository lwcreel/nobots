package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []user{
	{
		id:        "1",
		Name:      "John Doe",
		Username:  "jdog",
		Email:     "johndoe@example.com",
		Password:  "Password1",
		Followers: make([]user, 0),
		Following: make([]user, 0),
		Posts:     make([]post, 0),
		Comments:  make([]comment, 0),
	},
	{
		id:        "2",
		Name:      "Jane Doe",
		Username:  "jkitty",
		Email:     "janedoe@example.com",
		Password:  "Password1",
		Followers: make([]user, 0),
		Following: make([]user, 0),
		Posts:     make([]post, 0),
		Comments:  make([]comment, 0),
	},
	{
		id:        "3",
		Name:      "Adam Smith",
		Username:  "adamsapple24",
		Email:     "asmith23@example.com",
		Password:  "Password1",
		Followers: make([]user, 0),
		Following: make([]user, 0),
		Posts:     make([]post, 0),
		Comments:  make([]comment, 0),
	},
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	router.GET("/users", getUsers)

	// Start Server on Port 8080 (default)
	// Server will listen on 0.0.0.0.8080 (localhost:8080 on Windows)
	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
