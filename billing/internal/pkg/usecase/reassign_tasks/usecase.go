package reassign_tasks

import (
	"context"
	"math/rand"

	"async-arch/billing/internal/pkg/domain"
	"async-arch/billing/internal/pkg/producer"
)

type usecase struct {
	tasksRepository      domain.TasksRepository
	usersRepository      domain.UserRepository
	taskAssignedProducer producer.TaskAssignedProducer
}

func New(
	tasksRepository domain.TasksRepository,
	usersRepository domain.UserRepository,
	taskAssignedProducer producer.TaskAssignedProducer,
) *usecase {
	return &usecase{
		tasksRepository:      tasksRepository,
		usersRepository:      usersRepository,
		taskAssignedProducer: taskAssignedProducer,
	}
}

func (u *usecase) Run(ctx context.Context) error {
	tasks, err := u.tasksRepository.AllTasks(ctx)
	if err != nil {
		return err
	}

	employeeIDs, err := u.usersRepository.AllEmployeeIDs(ctx)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		randomIndex := rand.Intn(len(employeeIDs))
		pick := employeeIDs[randomIndex]

		task.AssigneeID = pick
	}

	tasks, err = u.tasksRepository.Upsert(ctx, tasks...)
	if err != nil {
		return err
	}

	return u.taskAssignedProducer.Produce(ctx, tasks...)
}
