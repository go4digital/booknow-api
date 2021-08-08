package dao

import (
	"github.com/go-pg/pg/v10"
	"github.com/go4digital/booknow-api/models"
)

var (
	Leads leadsInterface = &leads{}
)

const (
	LEADS_ALL = `SELECT id, firstname, lastname, email, phone, description FROM leads`
	LEADS     = `SELECT id, firstname, lastname, email, phone, description FROM leads WHERE id=$1`
	INSERT    = `insert into "leads"("firstname", "lastname", "email", "phone", "description") values($1, $2, $3, $4, $5) RETURNING id`
	UPDATE    = `UPDATE leads SET firstname=$2, lastname=$3, email=$4, phone=$5, description=$6 WHERE id=$1`
	DELETE    = `DELETE FROM leads WHERE id=$1`
)

type leadsInterface interface {
	CreateLead(*models.Lead) (int64, error)
	UpdateLead(*models.Lead) (int64, error)
	GetAllLeads() (*[]models.Lead, error)
	GetLead(int64) (*models.Lead, error)
	DeleteLead(int64) (int64, error)
}

type leads struct {
	DB *pg.DB
}

func NewLeads(db *pg.DB) leadsInterface {
	return &leads{DB: db}
}

func (leads *leads) CreateLead(lead *models.Lead) (int64, error) {

	var id int64

	err := leads.db.QueryRow(INSERT, lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Description).Scan(&id)

	checkNPrintError(err)

	return id, err
}

func (leads *leads) UpdateLead(lead *models.Lead) (int64, error) {

	res, err := leads.db.Exec(UPDATE, lead.ID, lead.FirstName, lead.LastName, lead.Email, lead.Phone, lead.Description)

	checkNPrintError(err)

	rowsAffected, err := res.RowsAffected()

	checkNPrintError(err)

	return rowsAffected, err
}

func (leads *leads) GetAllLeads() (*[]models.Lead, error) {

	var leadsArray []models.Lead

	rows, err := leads.db.Query(LEADS_ALL)

	checkNPrintError(err)

	// close the statement
	defer rows.Close()

	var lead models.Lead
	//iterate over the rows
	for rows.Next() {
		err = rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Description)
		checkNPrintError(err)
		leadsArray = append(leadsArray, lead)
	}

	return &leadsArray, err
}

func (leads *leads) GetLead(leadId int64) (*models.Lead, error) {

	var lead models.Lead

	rows := leads.db.QueryRow(LEADS, leadId)

	err := rows.Scan(&lead.ID, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &lead.Description)

	return &lead, err
}

func (leads *leads) DeleteLead(leadId int64) (int64, error) {

	res, err := leads.db.Exec(DELETE, leadId)

	checkNPrintError(err)

	rowsAffected, err := res.RowsAffected()

	checkNPrintError(err)

	return rowsAffected, err
}

func checkNPrintError(err error) {
	if err != nil {
		log.Error(err)
	}
}
