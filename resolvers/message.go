package resolvers

import (
	"context"

	"github.com/go4digital/booknow-api/models"
)

func (resolver *queryResolver) Message(ctx context.Context, id int) (*models.Message, error) {

	message, err := resolver.Service.GetMessage(id)

	return message, err
}
func (resolver *queryResolver) Messages(ctx context.Context) ([]models.Message, error) {

	messages, err := resolver.Service.GetAllMessages()

	return messages, err
}

func (resolver *mutationResolver) SaveMessage(ctx context.Context, message models.Message) (*models.Message, error) {

	messages, err := resolver.Service.SaveMessage(&message)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
