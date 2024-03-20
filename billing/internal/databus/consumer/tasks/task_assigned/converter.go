package task_assigned

import (
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"async-arch/billing/internal/databus/consumer/tasks"
)

func dtoToMsg(message kafka.Message) (*tasks.TaskAssignedMessage, error) {
	msg := &tasks.TaskAssignedMessage{}

	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
