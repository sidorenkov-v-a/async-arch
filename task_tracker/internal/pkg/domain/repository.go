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
}

type TasksRepository interface {
	Upsert(ctx context.Context, user *Task) (*Task, error)
}
