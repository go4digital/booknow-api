package services

import (
	"github.com/go4digital/booknow-api/dao"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
)

type Leads interface {
	CreateLead(*models.Lead) (*models.Lead, error)
	UpdateLead(*models.Lead) error
	GetAllLeads() ([]models.Lead, error)
	GetLead(int) (*models.Lead, error)
	DeleteLead(int) error
}

type leads struct {
	leadsDao dao.Leads
}

func NewLeads(leadsDao dao.Leads) Leads {
	return &leads{leadsDao: leadsDao}
}

func (service *leads) CreateLead(lead *models.Lead) (*models.Lead, error) {
	leadId, err := service.leadsDao.Create(lead)
	if err != nil {
		log.Error(err)
	}
	return leadId, err

}

func (service *leads) UpdateLead(lead *models.Lead) error {
	err := service.leadsDao.Update(lead)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (service *leads) GetAllLeads() ([]models.Lead, error) {
	leads, err := service.leadsDao.GetAll()
	if err != nil {
		log.Error(err)
	}
	return leads, err
}

func (service *leads) GetLead(id int) (*models.Lead, error) {
	lead, err := service.leadsDao.Get(id)
	if err != nil {
		log.Error(err)
	}
	return lead, err
}

func (service *leads) DeleteLead(id int) error {
	err := service.leadsDao.Delete(id)
	if err != nil {
		log.Error(err)
	}
	return err
}
