package reassign_tasks

import (
	"context"
)

type Usecase interface {
	Run(ctx context.Context) error
}
