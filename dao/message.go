package dao

import (
	"context"

	"github.com/go4digital/booknow-api/database"
	"github.com/go4digital/booknow-api/global"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
	"github.com/uptrace/bun"
)

type Message interface {
	Create(*models.Message) (*models.Message, error)
	Update(*models.Message) error
	GetAll() ([]models.Message, error)
	Get(int) (*models.Message, error)
	Delete(int) error
}
type message struct {
	db     *bun.DB
	ctx    context.Context
	cancel context.CancelFunc
}

func NewMessage(db *bun.DB) Message {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	return &message{db: db, ctx: ctx, cancel: cancel}
}

func (message *message) Create(input *models.Message) (*models.Message, error) {

	person := database.Person{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		ReferenceId: global.REFERENCES_ANONYMOUS,
	}

	_, err := message.db.NewInsert().Model(&person).Exec(message.ctx)
	checkNPrintError(err)

	contacts := []*database.Contact{
		{Description: input.Email, ReferenceId: global.REFERENCES_EMAIL},
		{Description: input.Phone, ReferenceId: global.REFERENCES_PHONE},
		{Description: input.Address, ReferenceId: global.REFERENCES_ADDRESS},
	}
	_, err = message.db.NewInsert().Model(&contacts).Exec(message.ctx)
	checkNPrintError(err)

	personContacts := []*database.PersonContact{}

	for _, contact := range contacts {
		personContacts = append(personContacts, &database.PersonContact{PersonId: person.Id, ContactId: contact.Id})
	}

	_, err = message.db.NewInsert().Model(&personContacts).Exec(message.ctx)
	checkNPrintError(err)

	responseMessage := database.Message{
		Description:  input.Description,
		FromPersonId: person.Id,
		ToPersonId:   global.BLOSSOMCLEAN_TENANT_ID,
		ReferenceId:  global.REFERENCES_ENQUIRY,
		CreatedBy:    global.REFERENCES_ANONYMOUS,
	}
	_, err = message.db.NewInsert().Model(&responseMessage).Exec(message.ctx)
	checkNPrintError(err)

	input.ID = responseMessage.Id

	return input, err
}

func (message *message) Update(input *models.Message) error {

	responseMessage := models.Message{
		ID:          input.ID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Phone:       input.Phone,
		Description: input.Description,
	}

	_, err := message.db.NewUpdate().Model(responseMessage).WherePK().Exec(message.ctx)

	checkNPrintError(err)

	return err
}

func (message *message) GetAll() ([]models.Message, error) {

	var messagesArray []models.Message

	err := message.db.NewSelect().Model((*models.Message)(nil)).Scan(message.ctx, &messagesArray)
	message.cancel()

	return messagesArray, err
}

func (message *message) Get(id int) (*models.Message, error) {

	responseMessage := models.Message{ID: id}

	err := message.db.NewSelect().Model(responseMessage).WherePK().Scan(message.ctx, &message)

	checkNPrintError(err)

	return &responseMessage, err
}

func (message *message) Delete(id int) error {
	responseMessage := models.Message{ID: id}

	_, err := message.db.NewDelete().Model(responseMessage).WherePK().Exec(message.ctx)

	checkNPrintError(err)

	return err
}

func checkNPrintError(err error) {
	if err != nil {
		log.Error(err)
	}
}
