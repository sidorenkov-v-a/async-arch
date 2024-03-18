package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	oapiMiddleware "github.com/oapi-codegen/nethttp-middleware"

	api_client "async-arch/billing/api/generated"
	"async-arch/billing/internal/api"
	"async-arch/billing/internal/api/middleware"
	"async-arch/billing/internal/infrastructure/contract"
	"async-arch/billing/internal/infrastructure/di"
	"async-arch/billing/internal/pkg/producer/task_assigned"
	"async-arch/billing/internal/pkg/producer/task_created"
	"async-arch/billing/internal/pkg/repository"
	"async-arch/billing/internal/pkg/usecase/create_task"
	"async-arch/billing/internal/pkg/usecase/reassign_tasks"
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
	taskCreatedProducer := task_created.NewProducer(databus)
	taskAssignedProducer := task_assigned.NewProducer(databus)

	// Usecases
	createTaskUsecase := create_task.New(tasksRepo, usersRepo, taskCreatedProducer, taskAssignedProducer)
	reassignTasksUsecase := reassign_tasks.New(tasksRepo, usersRepo, taskAssignedProducer)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(env.JWT, usersRepo)

	// API
	swagger, err := api_client.GetSwagger()
	if err != nil {
		return err
	}

	swagger.Servers = nil

	server := api.NewServer(createTaskUsecase, reassignTasksUsecase)

	r := mux.NewRouter()

	r.Use(oapiMiddleware.OapiRequestValidator(swagger))
	r.Use(middleware.JSONMiddleware)
	r.Use(authMiddleware.Handle)

	api_client.HandlerFromMux(server, r)

	// Run API Server
	apiServer := di.NewAPIServer(&env.Server)

	return apiServer.Run(r)
}
