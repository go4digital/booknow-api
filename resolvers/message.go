package resolvers

import (
	"context"
	"fmt"

	"github.com/go4digital/booknow-api/models"
)

func (resolver *queryResolver) Message(ctx context.Context, id int64) (*models.Message, error) {

	message, err := resolver.MessageService.GetMessage(id)

	return message, err
}
func (resolver *queryResolver) Messages(ctx context.Context) ([]models.Message, error) {

	messages, err := resolver.MessageService.GetAllMessages()

	return messages, err
}

func (resolver *mutationResolver) SaveMessage(ctx context.Context, message models.Message) (*models.Message, error) {

	folder := fmt.Sprintf("%v_%v", message.FirstName, message.Email)

	resolver.FileUploadService.Upload(folder, message.Files)

	messages, err := resolver.MessageService.SaveMessage(&message)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (resolver *mutationResolver) SaveEnquiry(ctx context.Context, message models.Message) (*models.Message, error) {

	messages, err := resolver.MessageService.SaveMessage(&message)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
