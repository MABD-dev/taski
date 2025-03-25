package validator

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mabd-dev/taski/internal/domain/config"
	"github.com/mabd-dev/taski/internal/domain/models"
)

// Task, check if task number, name, description, status and project name are valid data
//
// also see @TaskName, @TaskDescription, @TaskStatus and @TaskProject
func Task(task models.Task) error {
	if task.Number <= 0 {
		return fmt.Errorf("Invalid task number %v\n", task.Number)
	}
	if err := TaskName(task.Name); err != nil {
		return err
	}
	if err := TaskDescription(task.Description); err != nil {
		return err
	}
	if err := TaskStatus(task.Status); err != nil {
		return err
	}
	if err := TaskProject(task.Project); err != nil {
		return err
	}

	return nil
}

// TaskName, checks if name does not exceed max allowed number of characters and not blank, also
// removing whitespace
func TaskName(value string) error {
	nameLen := utf8.RuneCountInString(strings.TrimSpace(value))
	if nameLen <= 0 {
		return errors.New("Task name must not be empty")
	}
	if nameLen > config.TaskNameMaxLen {
		return fmt.Errorf("Task name must be smaller than %v characters", config.TaskNameMaxLen+1)
	}
	return nil
}

// TaskDescription, checks if description does not exceed max allowed number of characters and not blank, also
// removing whitespace
func TaskDescription(value string) error {
	descriptionLen := utf8.RuneCountInString(strings.TrimSpace(value))

	if descriptionLen > config.TaskDescriptionMaxLen {
		return fmt.Errorf("Task description must be smaller than %v characters", config.TaskDescriptionMaxLen+1)
	}

	return nil
}

// TaskStatus, checks if status is valid number. One of the @models.TaskStatus options
// removing whitespace
func TaskStatus(value models.TaskStatus) error {
	if value != models.Todo && value != models.InProgress && value != models.Done {
		return errors.New("invalid status")
	}
	return nil
}

// TaskProject, checks if project name does not exceed max allowed number of characters and not blank, also
// removing whitespace
func TaskProject(value string) error {
	projectLen := utf8.RuneCountInString(strings.TrimSpace(value))

	if projectLen > config.TaskProjectMaxLen {
		return fmt.Errorf("Task project must be smaller than %v characters", config.TaskProjectMaxLen+1)
	}

	return nil
}
