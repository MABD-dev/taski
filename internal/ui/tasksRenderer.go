package ui

import (
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/aquasecurity/table"
	"github.com/mabd-dev/taski/internal/domain/models"
)

func RenderTable(tasks []models.Task) {
	table := table.New(os.Stdout)
	table.SetHeaders("#", "Name", "Description", "Status", "Creation Date")

	for _, task := range tasks {
		datetimeFormatted := formatDatetime(task.CreatedAt)
		table.AddRow(strconv.Itoa(task.Number), task.Name, task.Description, task.Status.ToString(), datetimeFormatted)
	}
	table.Render()
}

func RenderKanbanBoard(tasks []models.Task) {
	rawData := TasksToRawData(tasks)
	RenderRawData(rawData)
}

func RenderRawData(data [][]string) {
	table := table.New(os.Stdout)

	for i, row := range data {
		if i == 0 { // header
			table.SetHeaders(row...)
		} else {
			table.AddRow(row...)
		}
	}
	table.Render()
}

func TasksToRawData(tasks []models.Task) [][]string {
	output := [][]string{}

	taskToStatusMap := map[models.TaskStatus][]models.Task{}
	taskToStatusMap[models.Todo] = []models.Task{}
	taskToStatusMap[models.InProgress] = []models.Task{}
	taskToStatusMap[models.Done] = []models.Task{}

	for _, task := range tasks {
		taskToStatusMap[task.Status] = append(taskToStatusMap[task.Status], task)
	}

	maxNumberOfRows := 0
	for _, statusTasks := range taskToStatusMap {
		maxNumberOfRows = max(maxNumberOfRows, len(statusTasks))
	}

	output = append(output, []string{models.Todo.ToString(), models.InProgress.ToString(), models.Done.ToString()})

	for i := range maxNumberOfRows {
		todoTaskName := ""
		inProgressTaskName := ""
		doneTaskName := ""

		if len(taskToStatusMap[models.Todo]) > i {
			todoTaskName = formatTaskForKanbanBoard(taskToStatusMap[models.Todo][i])
		}

		if len(taskToStatusMap[models.InProgress]) > i {
			inProgressTaskName = formatTaskForKanbanBoard(taskToStatusMap[models.InProgress][i])
		}

		if len(taskToStatusMap[models.Done]) > i {
			doneTaskName = formatTaskForKanbanBoard(taskToStatusMap[models.Done][i])
		}

		output = append(output, []string{todoTaskName, inProgressTaskName, doneTaskName})
	}

	return output
}

func formatTaskForKanbanBoard(task models.Task) string {
	var sb strings.Builder

	sb.WriteString(strconv.Itoa(task.Number))
	sb.WriteString(". ")
	sb.WriteString(task.Name)
	if utf8.RuneCountInString(task.Description) > 0 {
		sb.WriteString("\n - ")
		sb.WriteString(task.Description)
	}
	return sb.String()
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
