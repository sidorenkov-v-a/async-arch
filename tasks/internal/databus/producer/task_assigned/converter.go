package task_assigned

import (
	contract "async-arch/tasks/internal/databus/producer"
	"async-arch/tasks/internal/pkg/domain"
)

func taskToMsg(task *domain.Task) contract.TaskAssignedMessage {
	return contract.TaskAssignedMessage{
		ID:         task.ID,
		AssigneeID: task.AssigneeID,
	}
}
