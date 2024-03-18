package task_assigned

import (
	"context"
	"encoding/json"

	"async-arch/task_tracker/internal/infrastructure/di"
	"async-arch/task_tracker/internal/pkg/domain"
	"async-arch/task_tracker/internal/pkg/producer/converter"
	"async-arch/task_tracker/pkg/databus"
)

type producer struct {
	*databus.Producer
}

func NewProducer(databus *databus.Databus) *producer {
	p := di.NewProducer(databus, "tasks.task_assigned")
	return &producer{Producer: p}
}

func (p *producer) Produce(ctx context.Context, tasks ...*domain.Task) error {
	msgs := make([]*databus.Message, 0, len(tasks))

	for _, task := range tasks {
		msg := converter.TaskToTaskAssignedMessage(task)

		out, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		msgs = append(msgs, &databus.Message{Payload: out})
	}

	return p.Producer.Produce(ctx, msgs...)
}
