package task_created

import (
	"context"
	"encoding/json"

	"async-arch/tasks/internal/infrastructure/di"
	"async-arch/tasks/internal/pkg/domain"
	"async-arch/tasks/pkg/databus"
)

type producer struct {
	*databus.Producer
}

func New(databus *databus.Databus) *producer {
	p := di.NewProducer(databus, "tasks.task_created")
	return &producer{Producer: p}
}

func (p *producer) Produce(ctx context.Context, tasks ...*domain.Task) error {
	msgs := make([]*databus.Message, 0, len(tasks))

	for _, task := range tasks {
		msg := taskToMsg(task)

		out, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		msgs = append(msgs, &databus.Message{Payload: out})
	}

	return p.Producer.Produce(ctx, msgs...)
}
