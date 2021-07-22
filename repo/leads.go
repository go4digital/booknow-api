package leads

import (
	"log"

	"github.com/go4digital/booknow-api/models"
	"github.com/go4digital/booknow-api/postgres"
)

func InsertLead(lead models.Lead) int64 {
	db := postgres.Connect()

	// close database
	defer db.Close()

	sqlStatement := `insert into "leads"("firstname", "lastname", "email", "phone", "comments") 
	values($1, $2, $3, $4, $5) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Comments).Scan(&id)

	if err != nil {
		panic(err)
	}

	return id
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

		err = rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Comments)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		leads = append(leads, lead)
	}

	return leads, err

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
