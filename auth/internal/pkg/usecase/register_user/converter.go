package register_user

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"async-arch/auth/internal/pkg/domain"
	"async-arch/auth/pkg/databus"
)

type UserCreatedMessage struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func userToMessage(user *domain.User) (*databus.Message, error) {
	out, err := json.Marshal(UserCreatedMessage{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return &databus.Message{Payload: out}, nil
}
