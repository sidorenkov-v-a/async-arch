package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	oapiMiddleware "github.com/oapi-codegen/nethttp-middleware"

	api_client "async-arch/tasks/api/generated"
	"async-arch/tasks/internal/api"
	"async-arch/tasks/internal/api/middleware"
	"async-arch/tasks/internal/databus/producer/task_assigned"
	"async-arch/tasks/internal/databus/producer/task_completed"
	"async-arch/tasks/internal/databus/producer/task_created"
	"async-arch/tasks/internal/infrastructure/contract"
	"async-arch/tasks/internal/infrastructure/di"
	"async-arch/tasks/internal/pkg/repository"
	"async-arch/tasks/internal/pkg/usecase/complete_task"
	"async-arch/tasks/internal/pkg/usecase/create_task"
	"async-arch/tasks/internal/pkg/usecase/reassign_tasks"
)

const (
	success = 0
	fail    = 1
)

func main() {
	var err error

	log := di.NewLogger()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	defer func() {
		exitCode := success

		if panicErr := recover(); panicErr != nil {
			log.Error(ctx, panicErr)

			exitCode = fail
		}

		if err != nil {
			log.Error(log.WithError(err), "service running")

			exitCode = fail
		}

		os.Exit(exitCode)
	}()

	err = run(ctx, log)
}

func run(ctx context.Context, log contract.Log) (err error) {
	// Dependencies
	env, err := di.NewEnv()
	if err != nil {
		return err
	}

	databus := di.NewDatabus(env.Databus, log)

	// Database
	db, err := di.NewDB(env.DB)
	if err != nil {
		return err
	}

	// Repositories
	usersRepo := repository.NewUsersRepository(db)
	tasksRepo := repository.NewTasksRepository(db)

	//Producers
	taskCreatedProducer := task_created.New(databus)
	taskAssignedProducer := task_assigned.New(databus)
	taskCompletedProducer := task_completed.New(databus)

	// Usecases
	createTaskUsecase := create_task.New(tasksRepo, usersRepo, taskCreatedProducer, taskAssignedProducer)
	reassignTasksUsecase := reassign_tasks.New(tasksRepo, usersRepo, taskAssignedProducer)
	completeTaskUsecase := complete_task.New(tasksRepo, taskCompletedProducer)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(env.JWT, usersRepo)

	// API
	swagger, err := api_client.GetSwagger()
	if err != nil {
		return err
	}

	swagger.Servers = nil

	server := api.NewServer(createTaskUsecase, reassignTasksUsecase, completeTaskUsecase)

	r := mux.NewRouter()

	r.Use(oapiMiddleware.OapiRequestValidator(swagger))
	r.Use(middleware.JSONMiddleware)
	r.Use(authMiddleware.Handle)

	api_client.HandlerFromMux(server, r)

	// Run API Server
	apiServer := di.NewAPIServer(&env.Server)

	return apiServer.Run(r)
}
