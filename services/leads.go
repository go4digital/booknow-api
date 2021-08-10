package services

import (
	"github.com/go4digital/booknow-api/dao"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
)

type Leads struct {
	leadDao dao.Leads
}

func NewLeads(lead dao.Leads) dao.Leads {
	return &Leads{leadDao: lead}
}

func (service *Leads) Create(lead *models.LeadInput) (*models.Lead, error) {
	leadId, err := service.leadDao.Create(lead)
	if err != nil {
		log.Error(err)
	}
	return leadId, err

}

func (service *Leads) Update(lead *models.LeadInput) error {
	err := service.leadDao.Update(lead)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (service *Leads) GetAll() ([]models.Lead, error) {
	leads, err := service.leadDao.GetAll()
	if err != nil {
		log.Error(err)
	}
	return leads, err
}

func (service *Leads) Get(id int) (*models.Lead, error) {
	lead, err := service.leadDao.Get(id)
	if err != nil {
		log.Error(err)
	}
	return lead, err
}

func (service *Leads) Delete(id int) error {
	err := service.leadDao.Delete(id)
	if err != nil {
		log.Error(err)
	}
	return err
}
