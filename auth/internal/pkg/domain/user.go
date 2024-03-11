package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	Role         string    `db:"role"`
	HashPassword string    `db:"hash_password"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
