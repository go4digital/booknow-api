package services

import (
	"github.com/go4digital/booknow-api/dao"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
)

type Messages interface {
	SaveMessage(*models.Message) (*models.Message, error)
	GetAllMessages() ([]models.Message, error)
	GetMessage(int64) (*models.Message, error)
}

type message struct {
	messageDao dao.Message
}

func NewMessage(messageDao dao.Message) Messages {
	return &message{messageDao: messageDao}
}

func (service *message) SaveMessage(message *models.Message) (*models.Message, error) {
	leadId, err := service.messageDao.Create(message)
	if err != nil {
		log.Error(err)
	}
	return leadId, err

}

func (service *message) UpdateMessage(message *models.Message) error {
	err := service.messageDao.Update(message)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (service *message) GetAllMessages() ([]models.Message, error) {
	leads, err := service.messageDao.GetAll()
	if err != nil {
		log.Error(err)
	}
	return leads, err
}

func (service *message) GetMessage(id int64) (*models.Message, error) {
	lead, err := service.messageDao.Get(id)
	if err != nil {
		log.Error(err)
	}
	return lead, err
}
