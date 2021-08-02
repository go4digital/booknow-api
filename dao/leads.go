package leads

import (
	"database/sql"
	"log"
)

type Lead struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

const (
	LEADS_ALL = `SELECT id, firstname, lastname, email, phone, description FROM leads`
	LEADS     = `SELECT id, firstname, lastname, email, phone, description FROM leads WHERE id=$1`
	INSERT    = `insert into "leads"("firstname", "lastname", "email", "phone", "description") values($1, $2, $3, $4, $5) RETURNING id`
	UPDATE    = `UPDATE leads SET firstname=$2, lastname=$3, email=$4, phone=$5, description=$6 WHERE id=$1`
	DELETE    = `DELETE FROM leads WHERE id=$1`
)

func (l *Lead) InsertLead(db *sql.DB) (int64, error) {

	var id int64

	err := db.QueryRow(INSERT, l.FirstName, l.LastName, l.Email, l.Phone, l.Description).Scan(&id)

	checkError(err)

	return id, err
}

func (l *Lead) UpdateLead(db *sql.DB) (int64, error) {

	res, err := db.Exec(UPDATE, l.ID, l.FirstName, l.LastName, l.Email, l.Phone, l.Description)

	checkError(err)

	rowsAffected, err := res.RowsAffected()

	checkError(err)

	return rowsAffected, err
}

func GetAllLeads(db *sql.DB) ([]Lead, error) {

	var leads []Lead

	rows, err := db.Query(LEADS_ALL)

	checkError(err)

	// close the statement
	defer rows.Close()

	//itrate over the rows

	for rows.Next() {
		var lead Lead

		err = rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Description)

		checkError(err)

		leads = append(leads, lead)
	}

	return leads, err

}

func GetLead(db *sql.DB, leadId int64) (Lead, error) {

	var lead Lead

	rows := db.QueryRow(LEADS, leadId)

	err := rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Description)

	return lead, err
}

func DeleteLead(db *sql.DB, leadId int64) (int64, error) {

	res, err := db.Exec(DELETE, leadId)

	checkError(err)

	rowsAffected, err := res.RowsAffected()

	checkError(err)

	return rowsAffected, err
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
