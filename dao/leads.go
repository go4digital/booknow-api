package dao

import (
	"time"

	"github.com/go-pg/pg/v10"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
)

type Leads interface {
	Create(*models.LeadInput) (*models.Lead, error)
	Update(*models.LeadInput) error
	GetAll() ([]models.Lead, error)
	Get(int) (*models.Lead, error)
	Delete(int) error
}

type leads struct {
	db *pg.DB
}

func NewLeads(db *pg.DB) Leads {
	return &leads{db: db}
}

func (leads *leads) Create(input *models.LeadInput) (*models.Lead, error) {

	lead := models.Lead{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Phone:       input.Phone,
		Description: input.Description,
		CreatedAt:   time.Now(),
	}
	_, err := leads.db.Model(&lead).Insert()

	checkNPrintError(err)

	return &lead, err
}

func (leads *leads) Update(input *models.LeadInput) error {

	lead := models.Lead{
		ID:          input.ID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Phone:       input.Phone,
		Description: input.Description,
	}

	_, err := leads.db.Model(lead).WherePK().Update()

	checkNPrintError(err)

	return err
}

func (leads *leads) GetAll() ([]models.Lead, error) {

	var leadsArray []models.Lead

	err := leads.db.Model(&leadsArray).Select()

	checkNPrintError(err)

	return leadsArray, err
}

func (leads *leads) Get(id int) (*models.Lead, error) {

	lead := models.Lead{ID: id}

	err := leads.db.Model(lead).WherePK().Select()

	checkNPrintError(err)

	return &lead, err
}

func (leads *leads) Delete(id int) error {
	lead := models.Lead{ID: id}

	_, err := leads.db.Model(lead).WherePK().Delete()

	checkNPrintError(err)

	return err
}

func checkNPrintError(err error) {
	if err != nil {
		log.Error(err)
	}
}
