package cmd

import (
	"slices"

	"github.com/mabd-dev/taski/internal/domain/models"
)

func stringArrayToTaskStatus(strs []string) ([]models.TaskStatus, error) {
	statuses := []models.TaskStatus{}
	for _, statusStr := range strs {
		status, err := models.TaskStatusStrToStatus(statusStr)
		if err != nil {
			return statuses, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

func filterByStatus(tasks []models.Task, statuses []models.TaskStatus) []models.Task {
	filteredTasks := []models.Task{}

	for _, task := range tasks {
		if slices.Contains(statuses, task.Status) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
