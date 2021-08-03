package database

import (
	"database/sql"
	"log"

	"github.com/go4digital/booknow-api/utils"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	connectionString := utils.Getenv("CONNECTION_STR")
	databaseName := utils.Getenv("DRIVER_NAME")

	if connectionString == "" {
		log.Fatalf("Error: Empty Connection String !")
	}

	db, err := sql.Open(databaseName, connectionString)

	if err != nil {
		panic(err)
	}

	return db
}
