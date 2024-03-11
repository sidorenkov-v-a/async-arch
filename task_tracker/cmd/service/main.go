package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	oapiMiddleware "github.com/oapi-codegen/nethttp-middleware"

	api_client "async-arch/task_tracker/api/generated"
	"async-arch/task_tracker/internal/api"
	"async-arch/task_tracker/internal/api/middleware"
	"async-arch/task_tracker/internal/infrastructure/contract"
	"async-arch/task_tracker/internal/infrastructure/di"
	"async-arch/task_tracker/internal/pkg/repository"
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

	//databus := di.NewDatabus(env.Databus, log)

	// Database
	db, err := di.NewDB(env.DB)
	if err != nil {
		return err
	}

	// Repositories
	usersRepo := repository.NewUsersRepository(db)

	// Usecases

	// Middleware
	authMeddleware := middleware.NewAuthMiddleware(env.JWT, usersRepo)

	// API
	swagger, err := api_client.GetSwagger()
	if err != nil {
		return err
	}

	swagger.Servers = nil

	server := api.NewServer()

	r := mux.NewRouter()

	r.Use(oapiMiddleware.OapiRequestValidator(swagger))
	r.Use(middleware.JSONMiddleware)
	r.Use(authMeddleware.Handle)

	api_client.HandlerFromMux(server, r)

	// Run API Server
	apiServer := di.NewAPIServer(&env.Server)

	return apiServer.Run(r)
}
