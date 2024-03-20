package task_assigned

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
	msg, err := dtoToMsg(message)
	if err != nil {
		return err
	}

	_, err = h.tasksRepo.GetByTaskIDAndAssigneeID(ctx, msg.ID, msg.AssigneeID)
	if err != nil {
		return err
	}

	return nil
}
