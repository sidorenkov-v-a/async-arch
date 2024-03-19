package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Upsert(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	AllEmployeeIDs(ctx context.Context) ([]uuid.UUID, error)
}

type TasksRepository interface {
	Upsert(ctx context.Context, tasks ...*Task) ([]*Task, error)
	AllTasks(ctx context.Context) ([]*Task, error)
	GetByTaskIDAndAssigneeID(ctx context.Context, taskID uuid.UUID, assigneeID uuid.UUID) (*Task, error)
}
