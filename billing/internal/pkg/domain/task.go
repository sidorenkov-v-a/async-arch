package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `db:"id"`
	ReporterID  uuid.UUID  `db:"reporter_id"`
	AssigneeID  uuid.UUID  `db:"assignee_id"`
	JiraID      int64      `db:"jira_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Status      TaskStatus `db:"status"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}

type TaskStatus string

const (
	TaskStatusAssigned  TaskStatus = "assigned"
	TaskStatusCompleted TaskStatus = "completed"
)
