package services

import (
	"github.com/go4digital/booknow-api/dao"
	log "github.com/go4digital/booknow-api/logger"
	"github.com/go4digital/booknow-api/models"
)

type ICompany interface {
	CreateCompany(*models.Company) (*models.Company, error)
}

type company struct {
	companyDao dao.ICompany
}

func NewCompany(companyDao dao.ICompany) ICompany {
	return &company{companyDao: companyDao}
}

func (service *company) CreateCompany(company *models.Company) (*models.Company, error) {
	result, err := service.companyDao.Get(company.Name)
	if result.Id != 0 {
		return result, err
	}
	result, err = service.companyDao.Create(company)
	if err != nil {
		log.Error(err)
	}
	return result, err
}
