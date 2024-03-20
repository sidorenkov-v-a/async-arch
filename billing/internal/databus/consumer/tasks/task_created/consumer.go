package task_created

import (
	"context"
	"math/rand"

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

	task.Cost = rand.Int63n(20-10) + 10

	_, err = h.tasksRepo.Upsert(ctx, task)
	if err != nil {
		return err
	}

	return nil
}
