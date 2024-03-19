package complete_task

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	Run(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) error
}
