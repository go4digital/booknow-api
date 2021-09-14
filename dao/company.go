package dao

import (
	"context"

	"github.com/go4digital/booknow-api/database"
	"github.com/go4digital/booknow-api/global"
	"github.com/go4digital/booknow-api/models"
	"github.com/uptrace/bun"
)

type ICompany interface {
	Create(*models.Company) (*models.Company, error)
	Get(name string) (*models.Company, error)
}
type company struct {
	db     *bun.DB
	ctx    context.Context
	cancel context.CancelFunc
}

func NewCompany(db *bun.DB) ICompany {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	return &company{db: db, ctx: ctx, cancel: cancel}
}

func (company *company) Create(input *models.Company) (*models.Company, error) {

	newCompany := database.Company{
		Name: input.Name,
	}

	_, err := company.db.NewInsert().Model(&newCompany).Exec(company.ctx)
	checkNPrintError(err)

	companyContacts := []*database.Contact{
		{Description: input.Email, ReferenceId: global.REFERENCES_EMAIL},
		{Description: input.Mobile, ReferenceId: global.REFERENCES_MOBILE},
		{Description: input.ResidentialAddress, ReferenceId: global.REFERENCES_RESIDENTIALADDRESS},
		{Description: input.WebsiteAddress, ReferenceId: global.REFERENCES_WEBSITEADDRESS},
	}
	if input.Facebook != "" {
		companyContacts = append(companyContacts, &database.Contact{Description: input.Facebook, ReferenceId: global.REFERENCES_FACEBOOK})
	}
	if input.LinkedIn != "" {
		companyContacts = append(companyContacts, &database.Contact{Description: input.LinkedIn, ReferenceId: global.REFERENCES_LINKEDIN})
	}
	if input.GitHub != "" {
		companyContacts = append(companyContacts, &database.Contact{Description: input.GitHub, ReferenceId: global.REFERENCES_GITHUB})
	}
	if input.Instagram != "" {
		companyContacts = append(companyContacts, &database.Contact{Description: input.Instagram, ReferenceId: global.REFERENCES_INSTAGRAM})
	}
	if input.Whatsapp != "" {
		companyContacts = append(companyContacts, &database.Contact{Description: input.Whatsapp, ReferenceId: global.REFERENCES_WHATSAPP})
	}
	if input.Landline != "" {
		companyContacts = append(companyContacts, &database.Contact{Description: input.Landline, ReferenceId: global.REFERENCES_LANDLINE})
	}

	_, err = company.db.NewInsert().Model(&companyContacts).Exec(company.ctx)
	checkNPrintError(err)

	companyContactsMapping := []*database.CompanyContact{}

	for _, contact := range companyContacts {
		companyContactsMapping = append(companyContactsMapping, &database.CompanyContact{CompanyId: newCompany.Id, ContactId: contact.Id})
	}

	_, err = company.db.NewInsert().Model(&companyContactsMapping).Exec(company.ctx)
	checkNPrintError(err)

	admin := database.Person{
		FirstName: input.ContactPersonFirstName,
		LastName:  input.ContactPersonLastName,
		CreatedBy: newCompany.Id,
	}

	_, err = company.db.NewInsert().Model(&admin).Exec(company.ctx)
	checkNPrintError(err)

	companyPersonMapping := database.CompanyPerson{CompanyId: newCompany.Id, PersonId: admin.Id}

	_, err = company.db.NewInsert().Model(&companyPersonMapping).Exec(company.ctx)
	checkNPrintError(err)

	admnContacts := []*database.Contact{
		{Description: input.ContactPersonEmail, ReferenceId: global.REFERENCES_EMAIL},
		{Description: input.ContactPersonPhone, ReferenceId: global.REFERENCES_MOBILE},
	}
	if input.ContactPersonAddress != "" {
		admnContacts = append(admnContacts, &database.Contact{Description: input.ContactPersonAddress, ReferenceId: global.REFERENCES_RESIDENTIALADDRESS})
	}
	_, err = company.db.NewInsert().Model(&admnContacts).Exec(company.ctx)
	checkNPrintError(err)

	adminContactsMapping := []*database.PersonContact{}

	for _, contact := range admnContacts {
		adminContactsMapping = append(adminContactsMapping, &database.PersonContact{PersonId: admin.Id, ContactId: contact.Id})
	}
	_, err = company.db.NewInsert().Model(&adminContactsMapping).Exec(company.ctx)
	checkNPrintError(err)

	input.Id = newCompany.Id

	return input, err
}

func (company *company) Get(name string) (*models.Company, error) {
	result := new(database.Company)
	err := company.db.NewSelect().Model(result).Where("name = ?", name).Limit(1).Scan(company.ctx)
	response := models.Company{Id: result.Id, Name: result.Name}
	return &response, err
}
