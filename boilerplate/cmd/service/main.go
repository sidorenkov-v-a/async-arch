package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"

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

	// Databus test

	// make a new reader that consumes from topic-A

	fmt.Println("databus test start")
	err = databusTest()
	if err != nil {
		fmt.Println("databus test stop error")
		return err
	}
	fmt.Println("databus test stop success")

	//fmt.Println("start reader")
	//
	//databusReader := kafka.NewReader(kafka.ReaderConfig{
	//	Brokers:  []string{"localhost:9092"},
	//	GroupID:  "consumer-group-id",
	//	Topic:    "topic-A",
	//	MaxBytes: 10e6, // 10MB
	//})
	//
	//go func() {
	//	for {
	//		fmt.Println("read message")
	//
	//		m, err := databusReader.ReadMessage(context.Background())
	//		if err != nil {
	//			break
	//		}
	//		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	//	}
	//}()
	//
	//if err = databusReader.Close(); err != nil {
	//	return err
	//}
	//
	//fmt.Println("start writer")
	//databusWriter := &kafka.Writer{
	//	Addr:     kafka.TCP("localhost:9092"),
	//	Topic:    "topic-A",
	//	Balancer: &kafka.LeastBytes{},
	//}
	//
	//go func() {
	//	for {
	//		fmt.Println("write message")
	//
	//		err := databusWriter.WriteMessages(context.Background(),
	//			kafka.Message{
	//				Key:   []byte("Key-A"),
	//				Value: []byte("Hello World!"),
	//			},
	//			kafka.Message{
	//				Key:   []byte("Key-B"),
	//				Value: []byte("One!"),
	//			},
	//			kafka.Message{
	//				Key:   []byte("Key-C"),
	//				Value: []byte("Two!"),
	//			},
	//		)
	//		if err != nil {
	//			fmt.Println("failed to write messages:", err)
	//			break
	//		}
	//
	//		time.Sleep(1 * time.Second)
	//	}
	//
	//}()
	//
	//if err := databusWriter.Close(); err != nil {
	//	return err
	//}

	// Run API Server
	apiServer := di.NewAPIServer(&config.APIServer)

	return apiServer.Run(r)
}

//func databusTest() error {
//	topic := "my-topic"
//	partition := 0
//
//	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
//	if err != nil {
//		fmt.Println("failed to dial leader:", err)
//	}
//
//	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
//	_, err = conn.WriteMessages(
//		kafka.Message{Value: []byte("one!")},
//		kafka.Message{Value: []byte("two!")},
//		kafka.Message{Value: []byte("three!")},
//	)
//	if err != nil {
//		fmt.Println("failed to write messages:", err)
//	}
//
//	if err := conn.Close(); err != nil {
//		fmt.Println("failed to close writer:", err)
//	}
//
//	//-----------
//
//	conn, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
//	if err != nil {
//		fmt.Println("failed to dial leader:", err)
//	}
//
//	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
//	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
//
//	b := make([]byte, 10e3) // 10KB max per message
//	for {
//		n, err := batch.Read(b)
//		if err != nil {
//			break
//		}
//		fmt.Println(string(b[:n]))
//	}
//
//	if err := batch.Close(); err != nil {
//		fmt.Println("failed to close batch:", err)
//	}
//
//	if err := conn.Close(); err != nil {
//		fmt.Println("failed to close connection:", err)
//	}
//
//	return nil
//}

//func databusTest() error {
//	w := &kafka.Writer{
//		Addr:     kafka.TCP("localhost:9092"),
//		Topic:    "topic-A",
//		Balancer: &kafka.LeastBytes{},
//	}
//
//	fmt.Println("Writer write message")
//	err := w.WriteMessages(context.Background(),
//		kafka.Message{
//			Key:   []byte("Key-A"),
//			Value: []byte("Hello World!"),
//		},
//		kafka.Message{
//			Key:   []byte("Key-B"),
//			Value: []byte("One!"),
//		},
//		kafka.Message{
//			Key:   []byte("Key-C"),
//			Value: []byte("Two!"),
//		},
//	)
//	if err != nil {
//		fmt.Println("failed to write messages:", err)
//	}
//
//	if err := w.Close(); err != nil {
//		fmt.Println("failed to close writer:", err)
//	}
//
//	fmt.Println("Writer write message: success")
//
//	// make a new reader that consumes from topic-A
//	r := kafka.NewReader(kafka.ReaderConfig{
//		Brokers:  []string{"localhost:9092"},
//		GroupID:  "consumer-group-id",
//		Topic:    "topic-A",
//		MaxBytes: 10e6, // 10MB
//	})
//
//	for {
//		fmt.Println("Reader read message")
//		m, err := r.ReadMessage(context.Background())
//		if err != nil {
//			break
//		}
//		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
//	}
//
//	if err := r.Close(); err != nil {
//		fmt.Println("failed to close reader:", err)
//	}
//
//	return nil
//}

func databusTest() error {

	ctx := context.Background()

	go func() {

		w := &kafka.Writer{
			Addr:     kafka.TCP("localhost:9092"),
			Topic:    "topic-A",
			Balancer: &kafka.LeastBytes{},
		}

		for {

			fmt.Println("Writer write message", time.Now())
			err := w.WriteMessages(ctx,
				kafka.Message{
					Key:   []byte("Key-A"),
					Value: []byte("Hello World!"),
				},
				kafka.Message{
					Key:   []byte("Key-B"),
					Value: []byte("One!"),
				},
				kafka.Message{
					Key:   []byte("Key-C"),
					Value: []byte("Two!"),
				},
			)
			if err != nil {
				fmt.Println("failed to write messages:", err)
				break
			}

			fmt.Println("Writer write message: success")

			time.Sleep(1 * time.Second)
		}

		if err := w.Close(); err != nil {
			fmt.Println("failed to close writer:", err)
		}
	}()

	//fmt.Println("Writer write message", time.Now())
	//err := w.WriteMessages(context.Background(),
	//	kafka.Message{
	//		Key:   []byte("Key-A"),
	//		Value: []byte("Hello World!"),
	//	},
	//	kafka.Message{
	//		Key:   []byte("Key-B"),
	//		Value: []byte("One!"),
	//	},
	//	kafka.Message{
	//		Key:   []byte("Key-C"),
	//		Value: []byte("Two!"),
	//	},
	//)
	//if err != nil {
	//	fmt.Println("failed to write messages:", err)
	//	//break
	//}
	//
	//time.Sleep(1 * time.Second)

	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "topic-A",
		MaxBytes: 10e6, // 10MB
	})

	for {

		fmt.Println("Reader read message", time.Now())
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Read message error", err)
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		fmt.Println("failed to close reader:", err)
	}

	return nil
}
