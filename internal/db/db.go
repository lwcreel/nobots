package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectionString() string {
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

	return connectionString
}

func PostgresConnect(connectionString string) *pgx.Conn {

	conn, err := pgx.Connect(
		context.Background(),
		connectionString,
	)

	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to connect to DB: %v\n", err)
		os.Exit(1)
	}

	return conn
}
