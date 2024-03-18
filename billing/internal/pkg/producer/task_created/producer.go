package task_created

import (
	"context"
	"encoding/json"

	"async-arch/billing/internal/infrastructure/di"
	"async-arch/billing/internal/pkg/domain"
	"async-arch/billing/internal/pkg/producer/converter"
	"async-arch/billing/pkg/databus"
)

type producer struct {
	*databus.Producer
}

func NewProducer(databus *databus.Databus) *producer {
	p := di.NewProducer(databus, "tasks.task_created")
	return &producer{Producer: p}
}

func (p *producer) Produce(ctx context.Context, tasks ...*domain.Task) error {
	msgs := make([]*databus.Message, 0, len(tasks))

	for _, task := range tasks {
		msg := converter.TaskToTaskCreatedMessage(task)

		out, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		msgs = append(msgs, &databus.Message{Payload: out})
	}

	return p.Producer.Produce(ctx, msgs...)
}