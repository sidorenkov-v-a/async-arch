package complete_task

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"async-arch/tasks/internal/databus/producer"
	"async-arch/tasks/internal/pkg/domain"
)

var ErrNotFound = errors.New("not found")

type usecase struct {
	tasksRepository       domain.TasksRepository
	taskCompletedProducer producer.TaskCompletedProducer
}

func New(
	tasksRepository domain.TasksRepository,
	taskCompletedProducer producer.TaskCompletedProducer,
) *usecase {
	return &usecase{
		tasksRepository:       tasksRepository,
		taskCompletedProducer: taskCompletedProducer,
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

	err = u.taskCompletedProducer.Produce(ctx, task)
	if err != nil {
		return err
	}

	return nil
}
