package services

import (
	"log"

	"github.com/go4digital/booknow-api/dao"
	"github.com/go4digital/booknow-api/database"
)

var (
	LeadsService leadServiceInterface = &leadService{}
)

type leadService struct{}

type leadServiceInterface interface {
	CreateLead(*dao.Lead) (int64, error)
	UpdateLead(*dao.Lead) (int64, error)
	GetAllLeads() (*[]dao.Lead, error)
	GetLead(int64) (*dao.Lead, error)
	DeleteLead(int64) (int64, error)
}

var db = database.Connect()
var leads = dao.NewLeads(db)

func (leadService *leadService) CreateLead(lead *dao.Lead) (int64, error) {
	leadId, err := leads.CreateLead(lead)
	if err != nil {
		log.Println(err)
	}
	return leadId, err

}

func (leadService *leadService) UpdateLead(lead *dao.Lead) (int64, error) {
	rowsAffected, err := leads.UpdateLead(lead)
	if err != nil {
		log.Println(err)
	}
	return rowsAffected, err
}

func (leadService *leadService) GetAllLeads() (*[]dao.Lead, error) {
	leads, err := leads.GetAllLeads()
	if err != nil {
		log.Println(err)
	}
	return leads, err
}

func (leadService *leadService) GetLead(leadId int64) (*dao.Lead, error) {
	lead, err := leads.GetLead(leadId)
	if err != nil {
		log.Println(err)
	}
	return lead, err
}

func (leadService *leadService) DeleteLead(leadId int64) (int64, error) {
	rowsAffected, err := leads.DeleteLead(leadId)
	if err != nil {
		log.Println(err)
	}
	return rowsAffected, err
}
