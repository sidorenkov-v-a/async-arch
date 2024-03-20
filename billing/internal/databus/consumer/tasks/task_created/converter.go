package task_created

import (
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"async-arch/billing/internal/databus/consumer/tasks"
	"async-arch/billing/internal/pkg/domain"
)

func messageToTask(message kafka.Message) (*domain.Task, error) {
	msg := tasks.TaskCreatedMessage{}

	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          msg.ID,
		ReporterID:  msg.ReporterID,
		AssigneeID:  msg.AssigneeID,
		JiraID:      msg.JiraID,
		Title:       msg.Title,
		Description: msg.Description,
		Status:      msg.Status,
		CreatedAt:   msg.CreatedAt,
		UpdatedAt:   msg.UpdatedAt,
	}, nil
}
