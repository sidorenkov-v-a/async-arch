package main

import (
	"context"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"

	"async-arch/billing/internal/databus/consumer/auth/user_created"
	"async-arch/billing/internal/databus/consumer/tasks/task_created"
	"async-arch/billing/internal/infrastructure/contract"
	"async-arch/billing/internal/infrastructure/di"
	"async-arch/billing/internal/pkg/repository"
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

	// Handlers
	userCreatedHandler := user_created.New(usersRepo)
	taskUpsertedHandler := task_created.New(tasksRepo)

	userCreatedConsumer := di.NewConsumer(
		databus,
		"auth.user_created",
		"billing",
		userCreatedHandler.Handle,
	)
	taskCreatedConsumer := di.NewConsumer(
		databus,
		"tasks.task_created",
		"billing",
		taskUpsertedHandler.Handle,
	)

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return userCreatedConsumer.Consume(gCtx)
	})

	g.Go(func() error {
		return taskCreatedConsumer.Consume(gCtx)
	})

	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}
