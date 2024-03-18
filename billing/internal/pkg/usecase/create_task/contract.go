package create_task

import (
	"context"

	"async-arch/billing/internal/pkg/domain"
)

type Usecase interface {
	Run(ctx context.Context, in In) (*domain.Task, error)
}
