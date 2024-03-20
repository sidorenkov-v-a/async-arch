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

type tasksRepo struct {
	db *sqlx.DB
}

func NewTasksRepository(db *sqlx.DB) *tasksRepo {
	return &tasksRepo{db: db}
}

func (r *tasksRepo) Upsert(ctx context.Context, tasks ...*domain.Task) ([]*domain.Task, error) {
	query := `INSERT INTO tasks(id, reporter_id, assignee_id, jira_id, title, description, status, cost, created_at, updated_at)
VALUES (:id, :reporter_id, :assignee_id, :jira_id, :title, :description, :status, :cost, :created_at, :updated_at)
ON CONFLICT (id)
    DO UPDATE SET reporter_id = excluded.reporter_id,
                  assignee_id = excluded.assignee_id,
                  jira_id     = excluded.jira_id,
                  title       = excluded.title,
                  description = excluded.description,
                  status      = excluded.status,
                  cost        = excluded.cost,
                  created_at  = excluded.created_at,
                  updated_at  = excluded.updated_at
where tasks.updated_at < excluded.updated_at

RETURNING *;`

	res, err := sqlx.NamedQueryContext(ctx, trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db), query, tasks)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	out := make([]*domain.Task, 0, len(tasks))
	for res.Next() {
		task := domain.Task{}
		err = res.StructScan(&task)
		if err != nil {
			return nil, err
		}

		out = append(out, &task)
	}

	return out, nil
}

func (r *tasksRepo) AllTasks(ctx context.Context) ([]*domain.Task, error) {
	query := `select * from tasks;`

	tasks := make([]*domain.Task, 0, 100)

	err := sqlx.SelectContext(ctx, trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db), &tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *tasksRepo) GetByTaskIDAndAssigneeID(
	ctx context.Context,
	taskID uuid.UUID,
	assigneeID uuid.UUID,
) (*domain.Task, error) {
	query := `SELECT * FROM tasks WHERE id = ? and tasks.assignee_id = ?;`

	task := domain.Task{}

	if err := trmsqlx.DefaultCtxGetter.DefaultTrOrDB(ctx, r.db).
		GetContext(ctx, &task, r.db.Rebind(query), taskID, assigneeID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrTaskNotFound
		}

		return nil, err
	}

	return &task, nil
}
