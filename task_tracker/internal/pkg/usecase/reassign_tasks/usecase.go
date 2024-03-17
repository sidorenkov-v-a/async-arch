package reassign_tasks

import (
	"context"
	"math/rand"

	"async-arch/task_tracker/internal/pkg/domain"
)

type usecase struct {
	tasksRepository domain.TasksRepository
	usersRepository domain.UserRepository
	//taskCreatedProducer  *databus.Producer
	//TaskAssignedProducer *databus.Producer
}

func New(
	tasksRepository domain.TasksRepository,
	usersRepository domain.UserRepository,
	// usersRepository domain.UserRepository,
	// databus *databus.Databus,
) *usecase {
	//taskCreatedProducer := di.NewProducer(databus, "tasks.task_assigned")
	//taskAssignedProducer := di.NewProducer(databus, "tasks.task_created")

	return &usecase{
		tasksRepository: tasksRepository,
		usersRepository: usersRepository,
		//usersRepository:      usersRepository,
		//taskCreatedProducer:  taskCreatedProducer,
		//TaskAssignedProducer: taskAssignedProducer,
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

	_, err = u.tasksRepository.Upsert(ctx, tasks...)
	if err != nil {
		return err
	}

	return err
}
