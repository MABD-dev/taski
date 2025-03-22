package renderer

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mabd-dev/taski/internal/models"
)

func RenderTable(tasks []models.Task) {
	table := table.New(os.Stdout)
	table.SetHeaders("#", "Name", "Description", "Status", "Created At")

	for _, task := range tasks {
		datetimeFormatted := formatDatetime(task.CreatedAt)
		table.AddRow(strconv.Itoa(task.Number), task.Name, task.Description, task.Status.ToString(), datetimeFormatted)
	}
	table.Render()
}

// Maybe, just maybe this is not needed. But who cres, it's nice
func RenderTableByStatus(tasks []models.Task) {
	statusToTasks := make(map[models.TaskStatus][]models.Task)
	statusToTasks[models.Todo] = []models.Task{}
	statusToTasks[models.InProgress] = []models.Task{}
	statusToTasks[models.Done] = []models.Task{}

	for _, task := range tasks {
		statusToTasks[task.Status] = append(statusToTasks[task.Status], task)
	}

	fmt.Printf("%v tasks\n", models.Todo.ToString())
	RenderTable(statusToTasks[models.Todo])

	fmt.Printf("\n%v tasks\n", models.InProgress.ToString())
	RenderTable(statusToTasks[models.InProgress])

	fmt.Printf("\n%v tasks\n", models.Done.ToString())
	RenderTable(statusToTasks[models.Done])
}

func formatDatetime(datetime time.Time) string {
	switch {
	case isToday(datetime):
		return "Today"
	case isYesterday(datetime):
		return "Yesterday"
	case isTomorrow(datetime):
		return "Tomorrow"
	default:
		return datetime.Format(time.RFC1123)
	}
}

func isToday(t time.Time) bool {
	now := time.Now()
	return now.Year() == t.Year() && now.Month() == t.Month() && now.Day() == t.Day()
}

func isYesterday(t time.Time) bool {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	return yesterday.Year() == t.Year() && yesterday.Month() == t.Month() && yesterday.Day() == t.Day()
}

func isTomorrow(t time.Time) bool {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	return tomorrow.Year() == t.Year() && tomorrow.Month() == t.Month() && tomorrow.Day() == t.Day()
}
