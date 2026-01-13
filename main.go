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

	// TODO: Switch to Connection Pool for Concurrency
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
	router.POST("/users", postUsers(conn))

	// Start Server on Port 8080 (default)
	// Server will listen on 0.0.0.0.8080 (localhost:8080 on Windows)
	router.Run("localhost:8080")
}

// TODO: Errors in HTTP Requests Shoulnd't Kill Server

func getUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var users []user
		query := `SELECT * FROM users;`

		rows, _ := conn.Query(context.Background(), query)
		users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[user])
		if err != nil {
			log.Fatal("Error Fetching Row: " + err.Error())
			os.Exit(1)
		}

		c.IndentedJSON(http.StatusOK, users)
	}
	return gin.HandlerFunc(fn)
}

func postUsers(conn *pgx.Conn) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newUser user

		if err := c.BindJSON(&newUser); err != nil {
			log.Fatal("Error Binding JSON: " + err.Error())
			os.Exit(1)
		}

		query := `INSERT INTO users (name, username, email, passhash) VALUES (@name, @username, @email, @passhash) ON CONFLICT DO NOTHING`
		args := pgx.NamedArgs{
			"name":     newUser.Name,
			"username": newUser.Username,
			"email":    newUser.Email,
			"passhash": newUser.Passhash,
		}

		_, err := conn.Query(context.Background(), query, args)
		if err != nil {
			log.Fatal("Error Fetching Row: " + err.Error())
			os.Exit(1)
		}

		c.IndentedJSON(http.StatusCreated, newUser)
	}
	return gin.HandlerFunc(fn)
}
