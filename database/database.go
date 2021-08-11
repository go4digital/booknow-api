package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go4digital/booknow-api/global"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
)

func Connect() *pg.DB {

	connectionString := global.Getenv("CONNECTION_STRING")

	if connectionString == "" {
		log.Warn("Empty Connection String!")
	}

	opts, err := pg.ParseURL(connectionString)
	if err != nil {
		log.Error(err)
	}
	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		panic("Database connection failed")
	}

	return db
}

func CreateSchema(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createLeadTable(opts, db);
	return nil
}

func createLeadTable (opts *orm.CreateTableOptions, db *pg.DB) error {
    createError := db.Model(&models.Lead{}).CreateTable(opts)
    if createError != nil {
        log.Error(createError)
        return createError
    }
    log.Info("Lead table created")
    return nil
}
