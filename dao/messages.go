package dao

import (
	"context"

	"github.com/go4digital/booknow-api/database"
	"github.com/go4digital/booknow-api/global"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
	"github.com/uptrace/bun"
)

type Messages interface {
	Create(*models.Message) (*models.Message, error)
	Update(*models.Message) error
	GetAll() ([]models.Message, error)
	Get(int) (*models.Message, error)
	Delete(int) error
}
type messages struct {
	db  *bun.DB
	ctx context.Context
}

func NewMessages(db *bun.DB) Messages {
	parentCtx := context.Background()
	ctx, _ := context.WithCancel(parentCtx)
	return &messages{db: db, ctx: ctx}
}

func (messages *messages) Create(input *models.Message) (*models.Message, error) {

	person := database.Person{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		ReferenceId: global.REFERENCES_ANONYMOUS,
	}

	_, err := messages.db.NewInsert().Model(&person).Exec(messages.ctx)
	checkNPrintError(err)

	contacts := []*database.Contact{
		{Description: input.Email, ReferencesId: global.REFERENCES_EMAIL},
		{Description: input.Phone, ReferencesId: global.REFERENCES_PHONE},
		{Description: input.Address, ReferencesId: global.REFERENCES_ADDRESS},
	}
	_, err = messages.db.NewInsert().Model(&contacts).Exec(messages.ctx)
	checkNPrintError(err)

	personContacts := []*database.PersonContact{
		{PersonId: person.Id, ContactId: contacts[0].Id},
		{PersonId: person.Id, ContactId: contacts[1].Id},
		{PersonId: person.Id, ContactId: contacts[2].Id},
	}
	_, err = messages.db.NewInsert().Model(&personContacts).Exec(messages.ctx)
	checkNPrintError(err)

	message := database.Message{
		Description:  input.Description,
		FromPersonId: person.Id,
		ToPersonId:   global.BLOSSOMCLEAN_TENANT_ID,
		ReferencesId: global.REFERENCES_ENQUIRY,
	}
	_, err = messages.db.NewInsert().Model(&message).Exec(messages.ctx)
	checkNPrintError(err)

	input.ID = message.Id

	return input, err
}

func (messages *messages) Update(input *models.Message) error {

	message := models.Message{
		ID:          input.ID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Phone:       input.Phone,
		Description: input.Description,
	}

	_, err := messages.db.NewUpdate().Model(message).WherePK().Exec(messages.ctx)

	checkNPrintError(err)

	return err
}

func (messages *messages) GetAll() ([]models.Message, error) {

	var messagesArray []models.Message

	err := messages.db.NewSelect().Model((*models.Message)(nil)).Scan(messages.ctx, &messagesArray)

	return messagesArray, err
}

func (messages *messages) Get(id int) (*models.Message, error) {

	message := models.Message{ID: id}

	err := messages.db.NewSelect().Model(message).WherePK().Scan(messages.ctx, &message)

	checkNPrintError(err)

	return &message, err
}

func (messages *messages) Delete(id int) error {
	message := models.Message{ID: id}

	_, err := messages.db.NewDelete().Model(message).WherePK().Exec(messages.ctx)

	checkNPrintError(err)

	return err
}

func checkNPrintError(err error) {
	if err != nil {
		log.Error(err)
	}
}
