package task_created

import (
	contract "async-arch/task_tracker/internal/databus/producer"
	"async-arch/task_tracker/internal/pkg/domain"
)

func taskToMsg(task *domain.Task) contract.TaskCreatedMessage {
	return contract.TaskCreatedMessage{
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
