package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"

	"async-arch/boilerplate/internal/infrastructure/contract"
	"async-arch/boilerplate/internal/infrastructure/di"
	databus2 "async-arch/boilerplate/internal/infrastructure/di/databus"
	"async-arch/boilerplate/internal/infrastructure/di/env"
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

	// Databus test

	// make a new reader that consumes from topic-A

	err = databusExample(ctx, env, log)
	if err != nil {
		return err
	}

	// Run API Server
	apiServer := di.NewAPIServer(&config.APIServer)

	return apiServer.Run(r)
}

func databusExample(ctx context.Context, env *env.Config, log contract.Log) error {
	databus := di.NewDatabus(env.Databus, log)

	producer1 := di.NewProducer(databus, "topic_test_1")
	producer2 := di.NewProducer(databus, "topic_test_2")

	var handler1 databus2.ConsumerHandler = func(ctx context.Context, message kafka.Message) error {
		fmt.Println("-------------------------")
		fmt.Println("Handler 1 receive message")
		fmt.Println(message)

		return nil
	}

	var handler2 databus2.ConsumerHandler = func(ctx context.Context, message kafka.Message) error {
		fmt.Println("-------------------------")
		fmt.Println("Handler 2 receive message")
		fmt.Println(message)

		return nil
	}

	consumer1 := di.NewConsumer(databus, "topic_test_1", "consumer_1", handler1)
	consumer2 := di.NewConsumer(databus, "topic_test_2", "consumer_2", handler2)

	go func() {
		for {

			time.Sleep(1 * time.Second)

			msgs1 := []*databus2.Message{
				{
					Key:     "key1",
					Payload: []byte(time.Now().String()),
				},
				{
					Key:     "key2",
					Payload: []byte(time.Now().String()),
				},
			}

			msgs2 := []*databus2.Message{
				{
					Key:     "key3",
					Payload: []byte(time.Now().String()),
				},
				{
					Key:     "key4",
					Payload: []byte(time.Now().String()),
				},
			}

			err := producer1.Produce(ctx, msgs1...)
			if err != nil {
				fmt.Println("Error", err)
			}

			err = producer2.Produce(ctx, msgs2...)
			if err != nil {
				fmt.Println("Error", err)
			}
		}
	}()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() (err error) {
		return consumer1.Consume(gCtx)
	})

	g.Go(func() (err error) {
		return consumer2.Consume(gCtx)
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
