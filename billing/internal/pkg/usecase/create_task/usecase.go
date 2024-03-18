package create_task

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"async-arch/billing/internal/pkg/domain"
	"async-arch/billing/internal/pkg/producer"
)

var (
	ErrReporterNotFound      = errors.New("reporter not found")
	ErrAssigneeNotFound      = errors.New("assignee not found")
	ErrIncorrectAssigneeRole = errors.New("incorrect assignee role")
)

type In struct {
	ReporterID  uuid.UUID         `db:"reporter_id"`
	AssigneeID  uuid.UUID         `db:"assignee_id"`
	JiraID      int64             `db:"jira_id"`
	Title       string            `db:"title"`
	Description string            `db:"description"`
	Status      domain.TaskStatus `db:"status"`
}

type usecase struct {
	tasksRepository      domain.TasksRepository
	usersRepository      domain.UserRepository
	taskCreatedProducer  producer.TaskCreatedProducer
	taskAssignedProducer producer.TaskAssignedProducer
}

func New(
	tasksRepository domain.TasksRepository,
	usersRepository domain.UserRepository,
	taskCreatedProducer producer.TaskCreatedProducer,
	taskAssignedProducer producer.TaskAssignedProducer,
) *usecase {
	return &usecase{
		tasksRepository:      tasksRepository,
		usersRepository:      usersRepository,
		taskCreatedProducer:  taskCreatedProducer,
		taskAssignedProducer: taskAssignedProducer,
	}
}

func (u *usecase) Run(ctx context.Context, in In) (*domain.Task, error) {
	err := u.validate(ctx, in)
	if err != nil {
		return nil, err
	}

	tasks, err := u.tasksRepository.Upsert(ctx, &domain.Task{
		ID:          uuid.New(),
		ReporterID:  in.ReporterID,
		AssigneeID:  in.AssigneeID,
		JiraID:      in.JiraID,
		Title:       in.Title,
		Description: in.Description,
		Status:      "new",
	})

	task := tasks[0]

	err = u.taskCreatedProducer.Produce(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *usecase) validate(ctx context.Context, in In) error {
	isReporterExists, err := u.usersRepository.Exists(ctx, in.ReporterID)
	if err != nil {
		return err
	}

	if !isReporterExists {
		return ErrReporterNotFound
	}

	isAssigneeExists, err := u.usersRepository.Exists(ctx, in.AssigneeID)
	if err != nil {
		return err
	}

	if !isAssigneeExists {
		return ErrAssigneeNotFound
	}

	assignee, err := u.usersRepository.GetByID(ctx, in.AssigneeID)
	if err != nil {
		return err
	}

	if assignee.Role != "employee" {
		return ErrIncorrectAssigneeRole
	}

	return nil
}
