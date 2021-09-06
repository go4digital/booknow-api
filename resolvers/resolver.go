package resolvers

//go:generate go run github.com/99designs/gqlgen --verbose

import (
	"github.com/go4digital/booknow-api/graph/generated"
	"github.com/go4digital/booknow-api/services"
)

type Resolver struct {
	Service services.Messages
}

func (resolver *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{resolver}
}
func (resolver *Resolver) Query() generated.QueryResolver {
	return &queryResolver{resolver}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
