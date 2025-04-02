package repos

import (
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/converter"
	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/validator"
)

var (
	InvalidTaskNumber = errors.New("taskNumber cannot be negative")
)

type TasksRepoStruct struct {
	db        db.Db
	validator validator.Validator
}

var (
	TasksRepo TasksRepoStruct
)

func CreateTasksRepo(db db.Db, validator validator.Validator) TasksRepoStruct {
	TasksRepo = TasksRepoStruct{
		db:        db,
		validator: validator,
	}
	return TasksRepo
}

// GetAll return all tasks saved in memory
// @Returns
//
//	all saved tasks in memory db
func (repo *TasksRepoStruct) GetAll() []models.Task {
	return repo.db.GetAll()
}

// Get searches for a task with @taskNumber
//
// @Returns:
//
//	task if found or nil
func (repo *TasksRepoStruct) Get(taskNumber int) *models.Task {
	return repo.db.Get(taskNumber)
}

// ListWithFilters takes @statusFilters slice, and return all tasks that has a status
// included in @statusFilters
func (repo *TasksRepoStruct) ListWithFilters(statusFilters []string) []models.Task {
	tasks := repo.db.GetAll()
	if len(statusFilters) != 0 {
		statuses, err := converter.StringArrayToTaskStatus(statusFilters)
		if err != nil {
			panic(err)
		}
		tasks = converter.FilterByStatus(tasks, statuses)
	}
	return tasks
}

// Add takes task details and add it to inMemory tasks slice. New task number will
// be current max taskNumber + 1
//
// Parameters:
//   - name: task name, trim spaces and validate it
//   - description: task description, trim spaces and validate it
//   - status: task status
//   - project
func (repo *TasksRepoStruct) Add(
	name string,
	description string,
	status models.TaskStatus,
	project string,
	tags []string,
) error {
	trimmedName := strings.TrimSpace(name)
	trimmedDescription := strings.TrimSpace(description)
	trimmedProject := strings.TrimSpace(project)

	if err := repo.validator.TaskName(trimmedName); err != nil {
		return err
	}

	if err := repo.validator.TaskStatus(status); err != nil {
		return err
	}

	if err := repo.validator.TaskDescription(trimmedDescription); err != nil {
		return err
	}

	if err := repo.validator.TaskProject(trimmedProject); err != nil {
		return err
	}

	// Only add tags that length > 0 chars
	validatedTags := []string{}
	for _, tag := range tags {
		trimmedTag := strings.TrimSpace(tag)
		if utf8.RuneCountInString(trimmedTag) > 0 {
			validatedTags = append(validatedTags, trimmedTag)
		}
	}

	return repo.db.Add(trimmedName, trimmedDescription, status, trimmedProject, validatedTags)
}

// Update takes new task data and task number (ignoring that new task already has a
// task taskNumber)
// Searches for a task with @taskNumber if found it will be updated with new @task
// else return error
//
// @Parameters:
//   - task: task Number, Name, Descriptio, Status and Project are validated
//
// @Returns:
//
//	update task and return nil. error if task not found based on @taskNumber
func (repo *TasksRepoStruct) Update(taskNumber int, task models.Task) error {
	if taskNumber < 0 {
		return InvalidTaskNumber
	}

	if err := repo.validator.Task(task); err != nil {
		return err
	}

	task.Number = taskNumber

	return repo.db.Update(taskNumber, task)
}

// Delete taks list of taskNumbers and delete them.
//
// @Returns:
//
//	If any of the tasks for found  (base on it's taskNumber) will trow an error,
//	else update data and return nil
func (repo *TasksRepoStruct) Delete(taskNumbers ...int) error {
	for _, number := range taskNumbers {
		if number < 0 {
			return InvalidTaskNumber
		}
	}
	return repo.db.Delete(taskNumbers...)
}
