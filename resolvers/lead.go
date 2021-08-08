package resolvers

import (
	"context"

	"github.com/go4digital/booknow-api/models"
)

func (resolver *queryResolver) Lead(ctx context.Context, id int) (*models.Lead, error) {

	lead, err := resolver.LeadDao.GetLead(id)

	return lead, err
}
func (resolver *queryResolver) Leads(ctx context.Context) ([]models.Lead, error) {

	leads, err := resolver.LeadDao.GetAllLeads()

	return leads, err
}

func (resolver *mutationResolver) CreateLead(ctx context.Context, lead models.LeadInput) (*models.Lead, error) {

	leads, err := resolver.LeadDao.CreateLead(&lead)
	if err != nil {
		return nil, err
	}

	return leads, nil
}
