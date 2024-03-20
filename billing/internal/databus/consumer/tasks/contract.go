package tasks

import (
	"time"

	"github.com/google/uuid"

	"async-arch/billing/internal/pkg/domain"
)

type TaskAssignedMessage struct {
	ID         uuid.UUID `json:"id"`
	AssigneeID uuid.UUID `json:"assignee_id"`
}

type TaskCompletedMessage struct {
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
