package ui

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/aquasecurity/table"
	"github.com/fatih/color"
	"github.com/mabd-dev/taski/internal/domain/models"
)

var (
	taskNumberTitleFgColor      = color.New(color.FgHiBlue)
	taskNameTitleFgColor        = color.New(color.FgHiGreen)
	taskDescriptionTitleFgColor = color.New(color.FgHiGreen)
	taskProjFgColor             = color.New(color.FgHiCyan)
)

// RenderTable represent @tasks in a table with columns:
//   - TaskNumber, Name, Descrirption, Status, Creation Date
func RenderTable(tasks []models.Task) {
	table := table.New(os.Stdout)
	table.SetHeaders("#", "Name", "Description", "Status", "Creation Date")

	for _, task := range tasks {
		datetimeFormatted := formatDatetime(task.CreatedAt)
		table.AddRow(strconv.Itoa(task.Number), task.Name, task.Description, task.Status.ToString(), datetimeFormatted)
	}
	table.Render()
}

// RenderKanbanBoard takes @tasks, convert them to raw data using @TasksToKanbanRawData
// then draw them in a kanban style table
func RenderKanbanBoard(tasks []models.Task) {
	rawData := TasksToKanbanRawData(tasks)
	RenderRawData(rawData)
}

// RenderRawData take @data as 2d slice (rows and columns) and render them as a table
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

// TasksToKanbanRawData takes @tasks and convert them to 2d slice with columns:
//   - Todo, InProgress, Done
//
// Use @formatTaskForKanbanBoard to convert a taskt to string representation on kanban board
func TasksToKanbanRawData(tasks []models.Task) [][]string {
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

	c := color.New(color.FgHiGreen)

	output = append(output, []string{
		c.Sprint(models.Todo.ToString()),
		c.Sprint(models.InProgress.ToString()),
		c.Sprint(models.Done.ToString()),
	})

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

// RenderTask takes a task and diplay detailed data about the task
//
// Data displayed:
//   - task number
//   - task name
//   - task description
//   - task project: nicely formatted using @formatTaskProject
func RenderTask(task models.Task) {
	var sb strings.Builder

	// header
	sb.WriteString("╭")
	sb.WriteString(strings.Repeat("─", 50))
	sb.WriteString("\n")

	// task number
	sb.WriteString(taskNumberTitleFgColor.Sprintf("No. %v", task.Number))
	sb.WriteString("\n")

	// task name
	sb.WriteString(taskNameTitleFgColor.Sprint("Name: "))
	sb.WriteString(task.Name)
	sb.WriteString("\n")

	// task description
	sb.WriteString(taskDescriptionTitleFgColor.Sprint("Description: "))
	sb.WriteString(task.Description)
	sb.WriteString("\n")

	// task project
	sb.WriteString(formatTaskProject(task.Project))

	s := sb.String()
	re := regexp.MustCompile("\n")

	output := re.ReplaceAllStringFunc(s, func(match string) string {
		return match + "│"
	})
	fmt.Println(output)

	// footer
	fmt.Printf("╰%v\n", strings.Repeat("─", 50))
}

// formatTaskForKanbanBoard takes a task and format it's data to be displayed in
// kanban board
//
// Data displayed:
//   - task number
//   - task name: chunked into size of ~ 30 words see @chunkString
//   - task description: show '...' if description is not blank to tell user that their is a description
//   - task project name: colored and formatted nicely using @formatTaskProject
func formatTaskForKanbanBoard(task models.Task) string {
	var sb strings.Builder

	sb.WriteString(strconv.Itoa(task.Number))
	sb.WriteString(". ")

	sb.WriteString(chunkString(task.Name, 30))

	if utf8.RuneCountInString(task.Description) > 0 {
		sb.WriteString("\n...")
	}

	if utf8.RuneCountInString(task.Project) > 0 {
		sb.WriteString("\n")
		sb.WriteString(formatTaskProject(task.Project))
	}
	return sb.String()
}

// formatTaskProject takes task project name and add coloring to it, if it's not blank text
//
// Returns:
//
//	A string with color and styling for task project name
func formatTaskProject(value string) string {
	if utf8.RuneCountInString(strings.TrimSpace(value)) > 0 {
		return taskProjFgColor.Sprintf("@%v", value)
	}
	return value
}

// formatDatetime takes datetime and format in a nice way to be printed on the screen
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

// chunkSize takes a string and a chunk size, and it spluts the strings into n words
// (splitted by whitespace) it then concat words until their rune count is greater than
// or equal to chunkSize.
//
// Parameters:
//
//	s: The input string to be chunked.
//	chunkSize: The maximum rune count for each line.
//
// Returns:
//
//	A string containing the input string chunked into lines, with newlines inserted.
//
// Example:
//
//	input := "This is a long string that needs to be chunked into lines."
//	chunked := chunkString(input, 20)
//	// chunked will be:
//	// "This is a long string\nthat needs to be\nchunked into lines."
func chunkString(s string, chunkSize int) string {
	var sb strings.Builder

	chunks := strings.Split(s, " ")
	counter := 0

	for i, chunk := range chunks {
		counter += utf8.RuneCountInString(chunk)
		sb.WriteString(chunk)
		sb.WriteString(" ")

		if counter >= chunkSize && i != len(chunks)-1 {
			sb.WriteString("\n")
			counter = 0
		}
	}

	return sb.String()
}

// isToday checks if @t is today
func isToday(t time.Time) bool {
	now := time.Now()
	return now.Year() == t.Year() && now.Month() == t.Month() && now.Day() == t.Day()
}

// isYesterday checks if @t is yesterday
func isYesterday(t time.Time) bool {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	return yesterday.Year() == t.Year() && yesterday.Month() == t.Month() && yesterday.Day() == t.Day()
}

// isTomorrow checks if @t is tomorrow
func isTomorrow(t time.Time) bool {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	return tomorrow.Year() == t.Year() && tomorrow.Month() == t.Month() && tomorrow.Day() == t.Day()
}
