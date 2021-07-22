package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 2020
	user     = "BookNowAdmin"
	password = "BookNowAdmin@0987"
	dbname   = "BookNow"
)

func Connect() *sql.DB {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	return db
}
