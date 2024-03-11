package user_created

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"

	"async-arch/task_tracker/internal/pkg/domain"
)

type userCreatedMessage struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func messageToUser(message kafka.Message) (*domain.User, error) {
	userMessage := userCreatedMessage{}

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
