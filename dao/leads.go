package dao

import (
	"database/sql"
	"log"
)

var (
	LeadsDao leadsInterface = &leadsDao{}
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

type leadsInterface interface {
	CreateLead(*Lead) (int64, error)
	UpdateLead(*Lead) (int64, error)
	GetAllLeads() (*[]Lead, error)
	GetLead(int64) (*Lead, error)
	DeleteLead(int64) (int64, error)
}

type leadsDao struct {
	db *sql.DB
}

func NewLeadsDao(db *sql.DB) leadsInterface {
	return &leadsDao{db: db}
}

func (leadsDao *leadsDao) CreateLead(lead *Lead) (int64, error) {

	var id int64

	err := leadsDao.db.QueryRow(INSERT, lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Description).Scan(&id)

	checkError(err)

	return id, err
}

func (leadsDao *leadsDao) UpdateLead(lead *Lead) (int64, error) {

	res, err := leadsDao.db.Exec(UPDATE, lead.ID, lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Description)

	checkError(err)

	rowsAffected, err := res.RowsAffected()

	checkError(err)

	return rowsAffected, err
}

func (leadsDao *leadsDao) GetAllLeads() (*[]Lead, error) {

	var leads []Lead

	rows, err := leadsDao.db.Query(LEADS_ALL)

	checkError(err)

	// close the statement
	defer rows.Close()

	var lead Lead
	//iterate over the rows
	for rows.Next() {
		err = rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Description)
		checkError(err)
		leads = append(leads, lead)
	}

	return &leads, err
}

func (leadsDao *leadsDao) GetLead(leadId int64) (*Lead, error) {

	var lead Lead

	rows := leadsDao.db.QueryRow(LEADS, leadId)

	err := rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Description)

	return &lead, err
}

func (leadsDao *leadsDao) DeleteLead(leadId int64) (int64, error) {

	res, err := leadsDao.db.Exec(DELETE, leadId)

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
