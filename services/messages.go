package services

import (
	"github.com/go4digital/booknow-api/dao"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
)

type Messages interface {
	SaveMessage(*models.Message) (*models.Message, error)
	GetAllMessages() ([]models.Message, error)
	GetMessage(int) (*models.Message, error)
}

type messages struct {
	messagesDao dao.Messages
}

func NewMessages(messagesDao dao.Messages) Messages {
	return &messages{messagesDao: messagesDao}
}

func (service *messages) SaveMessage(message *models.Message) (*models.Message, error) {
	leadId, err := service.messagesDao.Create(message)
	if err != nil {
		log.Error(err)
	}
	return leadId, err

}

func (service *messages) UpdateMessage(message *models.Message) error {
	err := service.messagesDao.Update(message)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (service *messages) GetAllMessages() ([]models.Message, error) {
	leads, err := service.messagesDao.GetAll()
	if err != nil {
		log.Error(err)
	}
	return leads, err
}

func (service *messages) GetMessage(id int) (*models.Message, error) {
	lead, err := service.messagesDao.Get(id)
	if err != nil {
		log.Error(err)
	}
	return lead, err
}
