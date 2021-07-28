package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

// Connecting to db
func Connect() *pg.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := os.Getenv("CONNSTR")

	if connStr == "" {
		log.Fatalf("Error: Empty Connection String !")
	}
	opts, err := pg.ParseURL(connStr)
	if err != nil {
		log.Fatalf("Error: Parsing Connection String !")
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
