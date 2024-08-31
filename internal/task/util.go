package task

import "fmt"

func getIndexOfTask(tasks *[]Task, id uint64) (int, error) {
	for i, task := range *tasks {
		if task.ID == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("no task with id %v found", id)
}

func statusColor(status TaskStatus) string {
	switch status {
	case TASK_STATUS_TODO:
		return "8"
	case TASK_STATUS_IN_PROGRESS:
		return "#C26100"
	case TASK_STATUS_DONE:
		return "#00C274"
	default:
		return "8"
	}
}
