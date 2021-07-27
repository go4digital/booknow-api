package leads

import (
	"database/sql"
	"log"
)

type Lead struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Query     string `json:"query"`
}

func (l *Lead) InsertLead(db *sql.DB) int64 {

	sqlStatement := `insert into "leads"("firstname", "lastname", "email", "phone", "query") 
	values($1, $2, $3, $4, $5) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, l.FirstName, l.LastName, l.Email, l.Phone, l.Query).Scan(&id)

	checkError(err)

	return id
}

func (l *Lead) UpdateLead(db *sql.DB) int64 {

	sqlStatement := `UPDATE leads SET firstname=$2, lastname=$3, email=$4, phone=$5, query=$6 WHERE id=$1`

	res, err := db.Exec(sqlStatement, l.ID, l.FirstName, l.LastName, l.Email, l.Phone, l.Query)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected
}

func GetAllLeads(db *sql.DB) ([]Lead, error) {

	var leads []Lead

	sqlStatement := `SELECT * FROM leads`

	rows, err := db.Query(sqlStatement)

	checkError(err)

	// close the statement
	defer rows.Close()

	//itrate over the rows

	for rows.Next() {
		var lead Lead

		err = rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Query)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		leads = append(leads, lead)
	}

	return leads, err

}

func (l *Lead) GetLead(db *sql.DB) error {

	var lead Lead

	sqlStatement := `SELECT * FROM leads WHERE id=$1`

	rows := db.QueryRow(sqlStatement, l.ID)

	return rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Query)
}

func (l *Lead) DeleteLead(db *sql.DB) int64 {

	sqlStatement := `DELETE FROM leads WHERE id=$1`

	res, err := db.Exec(sqlStatement, l.ID)

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
