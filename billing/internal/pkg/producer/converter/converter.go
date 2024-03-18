package converter

import (
	"async-arch/billing/internal/pkg/domain"
	"async-arch/billing/internal/pkg/producer"
)

func TaskToTaskAssignedMessage(task *domain.Task) producer.TaskAssignedMessage {
	return producer.TaskAssignedMessage{
		ID:         task.ID,
		AssigneeID: task.AssigneeID,
	}
}

func TaskToTaskCreatedMessage(task *domain.Task) producer.TaskCreatedMessage {
	return producer.TaskCreatedMessage{
		ID:          task.ID,
		ReporterID:  task.ReporterID,
		AssigneeID:  task.AssigneeID,
		JiraID:      task.JiraID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
