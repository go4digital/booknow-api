package database

import (
	"database/sql"

	"github.com/go4digital/booknow-api/global"
	log "github.com/go4digital/booknow-api/logger"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	connectionString := global.Getenv("CONNECTION_STR")
	driverName := global.Getenv("DRIVER_NAME")

	if connectionString == "" {
		log.Warn("Error: Empty Connection String !")
	}

	db, err := sql.Open(driverName, connectionString)

	if err != nil {
	    log.Fatal(err)
		panic(err)
	}

	return db
}
