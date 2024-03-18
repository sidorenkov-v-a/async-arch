package databus

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"async-arch/billing/internal/infrastructure/contract"
)

type ConsumerHandler func(ctx context.Context, message kafka.Message) error

type Consumer struct {
	databus *Databus
	topic   string
	groupID string
	handler ConsumerHandler
	log     contract.Log
}

func NewConsumer(databus *Databus, topic string, groupID string, handler ConsumerHandler) *Consumer {
	return &Consumer{databus: databus, topic: topic, groupID: groupID, handler: handler, log: databus.log}
}

func (c *Consumer) Consume(ctx context.Context) (err error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{c.databus.broker},
		GroupID:  c.groupID,
		Topic:    c.topic,
		MaxBytes: 10e6, // 10MB
	})

	defer func() {
		if err != nil {
			c.log.Error(c.log.WithError(err), fmt.Sprintf("Databus: Consumer: failed to consume message: %s", c.topic))
		}

		err = reader.Close()
	}()

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		err = c.handler(ctx, message)
		if err != nil {
			c.log.Error(c.log.WithError(err), fmt.Sprintf("Databus: Consumer: failed to handle message: %s: skip", c.topic))
		}
	}
}
