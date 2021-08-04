package database

import (
	"database/sql"
	"log"

	"github.com/go4digital/booknow-api/global"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	connectionString := global.Getenv("CONNECTION_STR")
	drivename := global.Getenv("DRIVER_NAME")

	if connectionString == "" {
		log.Fatalf("Error: Empty Connection String !")
	}

	db, err := sql.Open(drivename, connectionString)

	if err != nil {
		panic(err)
	}

	return db
}
