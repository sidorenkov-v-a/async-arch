package repository

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"

	"async-arch/task_tracker/internal/pkg/domain"
)

type tasksRepo struct {
	db *sqlx.DB
}

func NewTasksRepository(db *sqlx.DB) *tasksRepo {
	return &tasksRepo{db: db}
}

func (r *tasksRepo) Upsert(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	query := `INSERT INTO tasks(id, reporter_id, assignee_id, title, description, status, updated_at)
VALUES (:id, :reporter_id, :assignee_id, :title, :description, :status, now())
ON CONFLICT (id)
    DO UPDATE SET reporter_id      = excluded.reporter_id,
                  assignee_id       = excluded.assignee_id,
                  title = excluded.title,
                  description  = excluded.description,
                  status = excluded.status,
                  updated_at = excluded.updated_at
where tasks.updated_at < excluded.updated_at

RETURNING *;`

	res, err := sqlx.NamedQueryContext(ctx, trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db), query, task)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	out := domain.Task{}
	if res.Next() {
		err = res.StructScan(&out)
		if err != nil {
			return nil, err
		}
	}

	return &out, nil
}
