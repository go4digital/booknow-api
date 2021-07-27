package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go4digital/booknow-api/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "BookNowAdmin",
		Password: "BookNowAdmin@0987",
		Addr:     "localhost:2020",
		Database: "BookNow",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateLeadTable(db)
	controllers.InitiateDB(db)
	return db
}
