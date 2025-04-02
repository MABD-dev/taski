package converter

import (
	"slices"

	"github.com/mabd-dev/taski/internal/domain/models"
)

// StringArrayToTaskStatus convert strs to slice of @models.TaskStatus
// if any str failed, return error
func StringArrayToTaskStatus(strs []string) ([]models.TaskStatus, error) {
	statuses := []models.TaskStatus{}
	for _, statusStr := range strs {
		status, err := models.TaskStatusStrToStatus(statusStr)
		if err != nil {
			return []models.TaskStatus{}, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

// FilterByStatus takes @tasks slice and @statuses slice, then filter
// tasks which has status in statuses slice
func FilterByStatus(tasks []models.Task, statuses []models.TaskStatus) []models.Task {
	filteredTasks := []models.Task{}

	for _, task := range tasks {
		if slices.Contains(statuses, task.Status) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
