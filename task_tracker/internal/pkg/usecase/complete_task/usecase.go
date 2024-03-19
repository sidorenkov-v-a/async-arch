package complete_task

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"async-arch/task_tracker/internal/pkg/domain"
)

var ErrNotFound = errors.New("not found")

type usecase struct {
	tasksRepository domain.TasksRepository
}

func New(
	tasksRepository domain.TasksRepository,
) *usecase {
	return &usecase{
		tasksRepository: tasksRepository,
	}
}

func (u *usecase) Run(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) error {
	task, err := u.tasksRepository.GetByTaskIDAndAssigneeID(ctx, taskID, userID)
	if err != nil {
		return err
	}

	task.Complete()

	_, err = u.tasksRepository.Upsert(ctx, task)
	if err != nil {
		return err
	}

	return nil
}
