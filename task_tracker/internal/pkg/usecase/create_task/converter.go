package create_task

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"async-arch/task_tracker/internal/pkg/domain"
	"async-arch/task_tracker/pkg/databus"
)

type TaskAssignedMessage struct {
	ID         uuid.UUID `json:"id"`
	AssigneeID uuid.UUID `json:"assignee_id"`
}

type TaskCreatedMessage struct {
	ID          uuid.UUID         `json:"id"`
	ReporterID  uuid.UUID         `json:"reporter_id"`
	AssigneeID  uuid.UUID         `json:"assignee_id"`
	JiraID      int64             `json:"jira_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      domain.TaskStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func taskToTaskAssignedMessage(task *domain.Task) (*databus.Message, error) {
	out, err := json.Marshal(TaskAssignedMessage{
		ID:         task.ID,
		AssigneeID: task.AssigneeID,
	})
	if err != nil {
		return nil, err
	}

	return &databus.Message{Payload: out}, nil
}

func taskToTaskCreatedMessage(task *domain.Task) (*databus.Message, error) {
	out, err := json.Marshal(TaskCreatedMessage{
		ID:          task.ID,
		ReporterID:  task.ReporterID,
		AssigneeID:  task.AssigneeID,
		JiraID:      task.JiraID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return &databus.Message{Payload: out}, nil
}
