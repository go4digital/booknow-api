package leads

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go4digital/booknow-api/models"
	"github.com/go4digital/booknow-api/postgres"
)

func InsertLead(lead models.Lead) int64 {
	db := postgres.Connect()

	// close database
	defer db.Close()

	sqlStatement := `insert into "leads"("firstname", "lastname", "email", "phone", "query") 
	values($1, $2, $3, $4, $5) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Query).Scan(&id)

	checkError(err)

	return id
}

func UpdateLead(lead models.Lead) int64 {
	db := postgres.Connect()

	// close database
	defer db.Close()

	sqlStatement := `UPDATE leads SET firstname=$2, lastname=$3, email=$4, phone=$5, query=$6 WHERE id=$1`

	res, err := db.Exec(sqlStatement, lead.ID, lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Query)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected
}

func GetAllLeads() ([]models.Lead, error) {
	db := postgres.Connect()

	// close database
	defer db.Close()

	var leads []models.Lead

	sqlStatement := `SELECT * FROM leads`

	rows, err := db.Query(sqlStatement)

	checkError(err)

	// close the statement
	defer rows.Close()

	//itrate over the rows

	for rows.Next() {
		var lead models.Lead

		err = rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Query)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		leads = append(leads, lead)
	}

	return leads, err

}

func GetLead(id int64) (models.Lead, error) {
	db := postgres.Connect()

	// close database
	defer db.Close()

	var lead models.Lead

	sqlStatement := `SELECT * FROM leads WHERE id=$1`

	rows := db.QueryRow(sqlStatement, id)

	err := rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Query)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return lead, nil
	case nil:
		return lead, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return lead, err

}

func DeleteLead(id int64) int64 {
	db := postgres.Connect()

	// close database
	defer db.Close()

	sqlStatement := `DELETE FROM leads WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
