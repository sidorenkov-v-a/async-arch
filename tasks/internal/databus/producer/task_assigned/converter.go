package task_assigned

import (
	contract "async-arch/task_tracker/internal/databus/producer"
	"async-arch/task_tracker/internal/pkg/domain"
)

func taskToMsg(task *domain.Task) contract.TaskAssignedMessage {
	return contract.TaskAssignedMessage{
		ID:         task.ID,
		AssigneeID: task.AssigneeID,
	}
}
