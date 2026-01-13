package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName

	conn, err := pgx.Connect(
		context.Background(),
		connectionString,
	)

	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to connect to DB: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	if err != nil {
		fmt.Fprint(os.Stderr, "QueryRow Failed: %v\n", err)
		os.Exit(1)
	}

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	router.GET("/users", getUsers(conn))
	router.POST("/users", postUsers)

	// Start Server on Port 8080 (default)
	// Server will listen on 0.0.0.0.8080 (localhost:8080 on Windows)
	router.Run("localhost:8080")
}

func getUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var row string
		err := conn.QueryRow(context.Background(), "SELECT name FROM users WHERE id=1;").Scan(&row)
		if err != nil {
			log.Fatal("Error Fetching Row: " + err.Error())
			os.Exit(1)
		}

		c.IndentedJSON(http.StatusOK, row)
	}
	return gin.HandlerFunc(fn)
}

func postUsers(c *gin.Context) {
	var newUser user

	// Binds the Received JSON to newUser
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Print(err)
		return
	}

	//users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
