package repl

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/mabd-dev/taski/internal/models"
	"github.com/mabd-dev/taski/internal/renderer"
)

// TODO: this code is redundent with cmd package, find a way to combine them
func list(s session, input string) error {
	tasks := s.db.List()

	// parsing status filtering flags
	statusValues := []string{}
	statusFunc := func(s string) error {
		statusValues = append(statusValues, s)
		return nil
	}

	listFlags := flag.NewFlagSet("list", flag.ContinueOnError)
	listFlags.Func("s", "Filter by status", statusFunc)

	parts := strings.Fields(input)
	listFlags.Parse(parts[1:])

	if len(statusValues) > 0 {
		statuses, err := stringArrayToTaskStatus(statusValues)
		if err != nil {
			return err
		}
		tasks = filterByStatus(tasks, statuses)
	}

	if len(tasks) == 0 {
		if len(statusValues) > 0 {
			fmt.Println("No tasks smatch this filter")
		} else {
			fmt.Println("You don't have any tasks yet")
		}
		return nil
	} else {
		renderer.RenderTable(tasks)
	}

	return nil
}

func help(s session, input string) error {
	fmt.Println("Welcome to taski REPL")
	fmt.Println("These are all the available commands")

	resetColor := "\033[0m"
	greenColor := "\033[32m"

	rawData := [][]string{}
	rawData = append(rawData, []string{"Command", "Description", "Alternatie Names"})

	for _, cmd := range getSortedCommands() {
		coloredName := greenColor + cmd.name + resetColor
		alternativeNames := strings.Join(cmd.alternativeNames, ", ")
		rawData = append(rawData, []string{coloredName, cmd.description, alternativeNames})
	}
	renderer.RenderRawData(rawData)

	return nil
}

func exit(s session, input string) error {
	os.Exit(0)
	return nil
}

func clear(s session, input string) error {
	// this is for linux only
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

// TODO: create a domain layer and put these function in it
// removes duplicate of these in cmd/helpers.js

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
