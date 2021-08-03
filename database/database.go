package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connectionString := os.Getenv("CONNECTION_STR")
	databaseName := os.Getenv("DRIVER_NAME")

	if connectionString == "" {
		log.Fatalf("Error: Empty Connection String !")
	}

	db, err := sql.Open(databaseName, connectionString)

	if err != nil {
		panic(err)
	}

	return db
}
