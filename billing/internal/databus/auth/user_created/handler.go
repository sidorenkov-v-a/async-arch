package user_created

import (
	"context"

	"github.com/segmentio/kafka-go"

	"async-arch/billing/internal/pkg/domain"
)

type handler struct {
	usersRepo domain.UserRepository
}

func New(usersRepo domain.UserRepository) *handler {
	return &handler{usersRepo: usersRepo}
}

func (h *handler) Handle(ctx context.Context, message kafka.Message) error {
	user, err := messageToUser(message)
	if err != nil {
		return err
	}

	_, err = h.usersRepo.Upsert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
