package resolvers

import (
	"context"
	"time"

	"github.com/go4digital/booknow-api/models"
)

func (resolver *queryResolver) Lead(ctx context.Context, id int) (*models.Lead, error) {
	lead := models.Lead{ID: id}

	if err := resolver.DB.Model(&lead).Select(); err != nil {
		return nil, err
	}

	return &lead, nil
}
func (resolver *queryResolver) Leads(ctx context.Context) ([]models.Lead, error) {
	var leads []models.Lead

	if err := resolver.DB.Model(&leads).Select(); err != nil {
		return nil, err
	}
	return leads, nil
}

func (resolver *mutationResolver) CreateLead(ctx context.Context, lead models.LeadInput) (*models.Lead, error) {
	l := models.Lead{
		FirstName:   lead.FirstName,
		LastName:    lead.LastName,
		Email:       lead.Email,
		Phone:       lead.Phone,
		Description: lead.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := resolver.DB.Model(&l).Insert()
	if err != nil {
		return nil, err
	}

	return &l, nil
}
