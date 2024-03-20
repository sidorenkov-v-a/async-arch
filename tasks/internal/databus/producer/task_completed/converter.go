package task_completed

import (
	contract "async-arch/tasks/internal/databus/producer"
	"async-arch/tasks/internal/pkg/domain"
)

func taskToMsg(task *domain.Task) contract.TaskCompletedMessage {
	return contract.TaskCompletedMessage{
		ID:         task.ID,
		AssigneeID: task.AssigneeID,
	}
}
