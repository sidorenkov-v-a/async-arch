package domain

import (
	"context"
)

type UserRepository interface {
	Upsert(ctx context.Context, in User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Exists(ctx context.Context, email string) (bool, error)
}
