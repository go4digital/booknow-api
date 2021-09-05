package dao

import (
	"context"
	"time"

	"github.com/go4digital/booknow-api/global"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
	"github.com/uptrace/bun"
)

type Leads interface {
	Create(*models.Lead) (*models.Lead, error)
	Update(*models.Lead) error
	GetAll() ([]models.Lead, error)
	Get(int) (*models.Lead, error)
	Delete(int) error
}
type leads struct {
	db  *bun.DB
	ctx context.Context
}

func NewLeads(db *bun.DB) Leads {
	parentCtx := context.Background()
	ctx, _ := context.WithCancel(parentCtx)
	return &leads{db: db, ctx: ctx}
}

func (leads *leads) Create(input *models.Lead) (*models.Lead, error) {

	lead := models.Lead{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Phone:       input.Phone,
		Description: input.Description,
		CreatedAt:   time.Now(),
	}

	person := models.Person{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		ReferenceId: global.REFERENCES_ANONYMOUS,
	}

	_, err := leads.db.NewInsert().Model(&person).Exec(leads.ctx)
	checkNPrintError(err)

	contacts := []*models.Contact{
		{Description: input.Email, ReferencesId: global.REFERENCES_EMAIL},
		{Description: input.Phone, ReferencesId: global.REFERENCES_PHONE},
	}
	_, err = leads.db.NewInsert().Model(&contacts).Exec(leads.ctx)
	checkNPrintError(err)

	personContacts := []*models.PersonContact{
		{PersonId: person.Id, ContactId: contacts[0].Id},
		{PersonId: person.Id, ContactId: contacts[1].Id},
	}
	_, err = leads.db.NewInsert().Model(&personContacts).Exec(leads.ctx)
	checkNPrintError(err)

	message := models.Message{
		Description:  input.Description,
		FromPersonId: person.Id,
		ToPersonId:   global.BLOSSOMCLEAN_TENANT_ID,
		ReferencesId: global.REFERENCES_ENQUIRY,
	}
	_, err = leads.db.NewInsert().Model(&message).Exec(leads.ctx)
	checkNPrintError(err)

	return &lead, err
}

func (leads *leads) Update(input *models.Lead) error {

	lead := models.Lead{
		ID:          input.ID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Phone:       input.Phone,
		Description: input.Description,
	}

	_, err := leads.db.NewUpdate().Model(lead).WherePK().Exec(leads.ctx)

	checkNPrintError(err)

	return err
}

func (leads *leads) GetAll() ([]models.Lead, error) {

	var leadsArray []models.Lead

	err := leads.db.NewSelect().Model((*models.Lead)(nil)).Scan(leads.ctx, &leadsArray)

	checkNPrintError(err)

	return leadsArray, err
}

func (leads *leads) Get(id int) (*models.Lead, error) {

	lead := models.Lead{ID: id}

	err := leads.db.NewSelect().Model(lead).WherePK().Scan(leads.ctx, &lead)

	checkNPrintError(err)

	return &lead, err
}

func (leads *leads) Delete(id int) error {
	lead := models.Lead{ID: id}

	_, err := leads.db.NewDelete().Model(lead).WherePK().Exec(leads.ctx)

	checkNPrintError(err)

	return err
}

func checkNPrintError(err error) {
	if err != nil {
		log.Error(err)
	}
}
