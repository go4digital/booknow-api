package database

import (
	"context"
	"database/sql"
	"os"

	log "github.com/go4digital/booknow-api/logger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect() *bun.DB {

	connectionString := os.Getenv("CONNECTION_STRING")

	if connectionString == "" {
		log.Warn("Empty Connection String!")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))

	db := bun.NewDB(sqldb, pgdialect.New())

	if db == nil {
		panic("Database connection failed")
	}

	return db
}

func CreateSchema(db *bun.DB) error {
	var err error
	ctx := context.Background()

	db.RegisterModel((*PersonContact)(nil))
	db.RegisterModel((*CompanyPerson)(nil))
	db.RegisterModel((*CompanyContact)(nil))

	models := []interface{}{
		(*Type)(nil),
		(*Reference)(nil),
		(*Person)(nil),
		(*Contact)(nil),
		(*PersonContact)(nil),
		(*Company)(nil),
		(*Document)(nil),
		(*CompanyContact)(nil),
		(*CompanyPerson)(nil),
		(*CompanyDocument)(nil),
		(*PersonDocument)(nil),
		(*Message)(nil),
	}
	for _, model := range models {
		if _, err = db.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
			return err
		}
	}

	return err
}
