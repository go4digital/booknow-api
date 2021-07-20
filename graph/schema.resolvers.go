package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go4digital/booknow-api/graph/generated"
	"github.com/go4digital/booknow-api/graph/model"
)

func (r *mutationResolver) CreateLead(ctx context.Context, input *model.LeadInput) (*model.Lead, error) {
	lead := &model.Lead{
		ID:        fmt.Sprintf("BNID%d", rand.Intn(100)),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
	}
	r.leads = append(r.leads, lead)
	return lead, nil
}

func (r *queryResolver) Leads(ctx context.Context) ([]*model.Lead, error) {
	return r.leads, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
