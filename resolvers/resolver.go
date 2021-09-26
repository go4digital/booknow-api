package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/go4digital/booknow-api/graph/generated"
	"github.com/go4digital/booknow-api/services"
)

type Resolver struct {
	MessageService    services.IMessages
	CompanyService    services.ICompany
	FileUploadService services.IFileUpload
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
