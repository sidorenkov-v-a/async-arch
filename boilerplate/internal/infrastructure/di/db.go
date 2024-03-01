package di

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"async-arch/boilerplate/internal/infrastructure/di/env"
)

func NewDB(env env.Database) (*sqlx.DB, error) {
	dbUrl := fmt.Sprintf(
		"postgresql://localhost:%d/postgres?user=%s&password=%s&sslmode=disable",
		env.DBPort,
		env.DBUser,
		env.DBPassword,
	)

	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
