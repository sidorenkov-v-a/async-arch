package task_created

import (
	"context"

	"github.com/segmentio/kafka-go"

	"async-arch/billing/internal/pkg/domain"
)

type handler struct {
	tasksRepo domain.TasksRepository
}

func New(tasksRepo domain.TasksRepository) *handler {
	return &handler{tasksRepo: tasksRepo}
}

func (h *handler) Handle(ctx context.Context, message kafka.Message) error {
	task, err := messageToTask(message)
	if err != nil {
		return err
	}

	_, err = h.tasksRepo.Upsert(ctx, task)
	if err != nil {
		return err
	}

	return nil
}
