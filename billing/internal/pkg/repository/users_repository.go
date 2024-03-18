package repository

import (
	"context"
	"database/sql"
	"errors"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"async-arch/billing/internal/pkg/domain"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *usersRepo {
	return &usersRepo{db: db}
}

func (r *usersRepo) Upsert(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users(id, email, role, first_name, last_name, created_at, updated_at)
VALUES (:id, :email, :role, :first_name, :last_name, :created_at, :updated_at)
ON CONFLICT (id)
    DO UPDATE SET email      = excluded.email,
                  role       = excluded.role,
                  first_name = excluded.first_name,
                  last_name  = excluded.last_name,
                  created_at = excluded.created_at,
                  updated_at = excluded.updated_at
where users.updated_at < excluded.updated_at

RETURNING *;`

	res, err := sqlx.NamedQueryContext(ctx, trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db), query, user)
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

func (r *usersRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)`

	out := false

	if err := trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db).
		GetContext(ctx, &out, r.db.Rebind(query), id); err != nil {

		return false, err
	}

	return out, nil
}

func (r *usersRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `SELECT * FROM users WHERE id = ?;`

	user := domain.User{}

	if err := trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db).
		GetContext(ctx, &user, r.db.Rebind(query), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *usersRepo) AllEmployeeIDs(ctx context.Context) ([]uuid.UUID, error) {
	query := `SELECT users.id FROM users where role = 'employee'`

	employeeIDs := make([]uuid.UUID, 0, 0)

	if err := trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db).
		SelectContext(ctx, &employeeIDs, r.db.Rebind(query)); err != nil {
		return nil, err
	}

	return employeeIDs, nil
}
