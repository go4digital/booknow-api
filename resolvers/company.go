package resolvers

import (
	"context"

	"github.com/go4digital/booknow-api/models"
)

func (resolver *mutationResolver) CreateCompany(ctx context.Context, input models.Company) (*models.Company, error) {

	company, err := resolver.CompanyService.CreateCompany(&input)
	if err != nil {
		return nil, err
	}

	return company, nil
}
