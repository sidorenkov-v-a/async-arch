package repository

import (
	"context"
	"database/sql"
	"errors"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"

	"async-arch/auth/internal/pkg/domain"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *usersRepo {
	return &usersRepo{db: db}
}

func (r *usersRepo) Upsert(ctx context.Context, in domain.User) (*domain.User, error) {
	query := `INSERT INTO users(id, email, role, hash_password, first_name, last_name, updated_at)
VALUES (:id, :email, :role, :hash_password, :first_name, :last_name, NOW())
ON CONFLICT (id)
    DO UPDATE SET email         = excluded.email,
                  role          = excluded.role,
                  hash_password = excluded.hash_password,
                  first_name    = excluded.first_name,
                  last_name     = excluded.last_name
RETURNING *;`

	res, err := sqlx.NamedQueryContext(ctx, trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db), query, in)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	out := domain.User{}
	if res.Next() {
		err = res.StructScan(&out)
		if err != nil {
			return nil, err
		}
	}

	return &out, nil
}

func (r *usersRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE email = ?;`

	user := domain.User{}

	if err := trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db).
		GetContext(ctx, &user, r.db.Rebind(query), email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *usersRepo) Exists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`

	out := false

	if err := trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db).
		GetContext(ctx, &out, r.db.Rebind(query), email); err != nil {

		return false, err
	}

	return out, nil
}
