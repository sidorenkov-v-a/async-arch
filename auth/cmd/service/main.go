package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	oapiMiddleware "github.com/oapi-codegen/nethttp-middleware"

	api_client "async-arch/auth/api/generated"
	"async-arch/auth/internal/api"
	"async-arch/auth/internal/api/middleware"
	"async-arch/auth/internal/infrastructure/contract"
	"async-arch/auth/internal/infrastructure/di"
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

	// Database
	//_, err = di.NewDB(env.DB)
	//if err != nil {
	//	return err
	//}

	// Repositories

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

	api_client.HandlerFromMux(server, r)

	// Run API Server
	apiServer := di.NewAPIServer(&env.Server)

	return apiServer.Run(r)
}
