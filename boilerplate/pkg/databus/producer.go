package databus

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"

	"async-arch/boilerplate/internal/infrastructure/contract"
)

type Message struct {
	Key     string
	Payload []byte
}

type Producer struct {
	databus *Databus
	topic   string
	log     contract.Log
}

func NewProducer(databus *Databus, topic string) *Producer {
	return &Producer{
		databus: databus,
		topic:   topic,
		log:     databus.log,
	}
}

func (p *Producer) Produce(ctx context.Context, msgs ...*Message) (err error) {
	writer := kafka.Writer{
		Addr:         kafka.TCP(p.databus.broker),
		Topic:        p.topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1,
		BatchTimeout: 0,
		WriteTimeout: 0,
	}

	defer func() {
		if err != nil {
			p.log.Error(p.log.WithError(err), "Databus: Producer: failed to produce messages")
		}

		err = writer.Close()
	}()
	fmt.Println("Write message", time.Now())

	kafkaMsgs := make([]kafka.Message, 0, len(msgs))

	for _, message := range msgs {
		kafkaMsgs = append(kafkaMsgs, kafka.Message{
			Key:   []byte(message.Key),
			Value: message.Payload,
		})
	}

	err = writer.WriteMessages(ctx, kafkaMsgs...)
	if err != nil {
		return err
	}

	return nil
}
