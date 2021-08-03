package database

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go4digital/booknow-api/utils"
)

// Connecting to db
func Connect() *pg.DB {

	connectionString := utils.Getenv("CONNECTION_STRING")

	if connectionString == "" {
		log.Fatalf("Error: Empty Connection String !")
	}
	opts, err := pg.ParseURL(connectionString)
	if err != nil {
		log.Fatalf("Error: Parsing Connection String !")
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Println("Failed to connect")
		os.Exit(100)
	}
	log.Println("Connected to database")
	return db
}
