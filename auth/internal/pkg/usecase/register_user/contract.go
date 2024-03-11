package register_user

import (
	"context"

	"async-arch/auth/internal/pkg/domain"
)

type Usecase interface {
	Run(ctx context.Context, in In) (*domain.User, error)
}
