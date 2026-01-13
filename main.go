package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []user{
	{
		id:        "1",
		Name:      "John Doe",
		Username:  "jdog",
		Email:     "johndoe@example.com",
		Password:  "",
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
		Password:  "",
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
		Password:  "",
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
	router.POST("/users", postUsers)

	// Start Server on Port 8080 (default)
	// Server will listen on 0.0.0.0.8080 (localhost:8080 on Windows)
	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser user

	// Binds the Received JSON to newUser
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Print(err)
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
