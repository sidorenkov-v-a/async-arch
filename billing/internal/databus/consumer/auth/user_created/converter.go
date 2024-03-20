package user_created

import (
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"async-arch/billing/internal/databus/consumer/auth"
	"async-arch/billing/internal/pkg/domain"
)

func messageToUser(message kafka.Message) (*domain.User, error) {
	userMessage := auth.UserCreatedMessage{}

	err := json.Unmarshal(message.Value, &userMessage)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        userMessage.ID,
		Email:     userMessage.Email,
		Role:      userMessage.Role,
		FirstName: userMessage.FirstName,
		LastName:  userMessage.LastName,
		CreatedAt: userMessage.CreatedAt,
		UpdatedAt: userMessage.UpdatedAt,
	}, nil
}
