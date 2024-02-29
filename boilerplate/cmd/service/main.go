package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	"async-arch/boilerplate/internal/infrastructure/contract"
	"async-arch/boilerplate/internal/infrastructure/di"
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

	config, err := di.NewConfig()
	if err != nil {
		return err
	}

	env, err := di.NewEnv()
	if err != nil {
		return err
	}

	// Database
	_, err = di.NewDB(env.DB)
	if err != nil {
		return err
	}

	// Repositories

	// API

	//TODO:
	//https://github.com/deepmap/oapi-codegen/tree/master/examples/petstore-expanded/gorilla
	r := mux.NewRouter()

	// Run API Server
	apiServer := di.NewAPIServer(&config.APIServer)

	return apiServer.Run(r)
}
