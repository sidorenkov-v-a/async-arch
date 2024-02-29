package di

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"async-arch/boilerplate/internal/infrastructure/di/env"
)

func NewDB(env env.Database) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", env.DBUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
