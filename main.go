package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lwcreel/nobots/internal/db"
	"github.com/lwcreel/nobots/internal/user"
)

func main() {

	// TODO: Switch to Connection Pool for Concurrency
	conn := db.PostgresConnect(db.ConnectionString())
	defer conn.Close(context.Background())

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.POST("/users", user.PostUsers(conn))
	router.GET("/users", user.GetUsers(conn))
	router.DELETE("/users", user.DeleteUsers(conn))
	router.PUT("/users", user.PutUsers(conn))

	// Start Server on Port 8080 (default)
	// Server will listen on 0.0.0.0.8080 (localhost:8080 on Windows)
	router.Run("localhost:8080")
}
