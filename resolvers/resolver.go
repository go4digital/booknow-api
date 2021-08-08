package resolvers

//go:generate go run github.com/99designs/gqlgen --verbose

import (
	"github.com/go4digital/booknow-api/dao"
	"github.com/go4digital/booknow-api/graph/generated"
)

type Resolver struct {
	LeadDao dao.LeadsInterface
}

func (resolver *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{resolver}
}
func (resolver *Resolver) Query() generated.QueryResolver {
	return &queryResolver{resolver}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
