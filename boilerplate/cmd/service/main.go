package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	"async-arch/boilerplate/internal/api"
	"async-arch/boilerplate/internal/api/handler/index"
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

	apiServer := di.NewAPIServer(&config.APIServer)

	// Database

	// Repositories

	// API
	r := mux.NewRouter()
	apiWrapper := api.NewWrapper(log)

	r.Handle(apiWrapper.Handle(index.New())).Methods(http.MethodGet)

	// Run API Server
	return apiServer.Run(r)
}
