package cmd

import (
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mabd-dev/tasks/internal/db"
	"github.com/mabd-dev/tasks/internal/models"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all your tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		statuses, err := cmd.Flags().GetStringArray("status")
		if err != nil {
			panic(err)
		}

		db := db.GetDb()
		tasks := db.List()

		if len(statuses) != 0 {
			statuses, err := stringArrayToTaskStatus(statuses)
			if err != nil {
				panic(err)
			}
			tasks = filterByStatus(tasks, statuses)
		}

		table := table.New(os.Stdout)
		table.SetHeaders("#", "Name", "Description", "Status", "Created At")

		for _, task := range tasks {
			table.AddRow(strconv.Itoa(task.Number), task.Name, task.Description, task.Status.ToString(), task.CreatedAt.Format(time.RFC1123))
		}
		table.Render()
	},
}

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
