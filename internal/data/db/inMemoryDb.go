package db

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/mabd-dev/taski/internal/domain/models"
)

type InMemoryDb struct {
	Tasks *[]models.Task
}

// GetAll return all tasks saved in memory
// @Returns
//
//	all saved tasks in memory db
func (db *InMemoryDb) GetAll() []models.Task {
	return *db.Tasks
}

// Add takes task details and add it to inMemory tasks slice. New task number will
// be current max taskNumber + 1
func (db *InMemoryDb) Add(
	name string,
	description string,
	status models.TaskStatus,
	project string,
) error {

	maxTaxNumber := 0
	for _, task := range *db.Tasks {
		maxTaxNumber = max(maxTaxNumber, task.Number)
	}
	newTaskNumber := maxTaxNumber + 1

	newTask := models.Task{
		Number:      newTaskNumber,
		Name:        name,
		Description: description,
		Status:      status,
		Project:     project,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	*db.Tasks = append(*db.Tasks, newTask)
	return nil
}

// Get searches for a task with @taskNumber
//
// @Returns:
//
//	task if found or nil
func (db *InMemoryDb) Get(taskNumber int) *models.Task {
	for _, t := range *db.Tasks {
		if t.Number == taskNumber {
			return &t
		}
	}
	return nil
}

// Update takes new task data and task number (ignoring that new task already has a
// task taskNumber)
// Searches for a task with @taskNumber if found it will be updated with new @task
// else return error
//
// @Returns:
//
//	update task and return nil. error if task not found based on @taskNumber
func (db *InMemoryDb) Update(taskNumber int, task models.Task) error {
	taskIndex := db.getTaskIndexFromNumber(taskNumber)
	if taskIndex == -1 {
		return errors.New("Could not find task")
	}

	tasks := *db.Tasks
	tasks[taskIndex] = task

	return nil
}

// Deelete taks list of taskNumbers and delete them.
//
// @Returns:
//
//	If any of the tasks for found  (base on it's taskNumber) will trow an error,
//	else update data and return nil
func (db *InMemoryDb) Delete(taskNumbers ...int) error {
	for _, taskNumber := range taskNumbers {
		taskIndex := db.getTaskIndexFromNumber(taskNumber)
		if taskIndex == -1 {
			return fmt.Errorf("Could not find task with number=%v", taskNumber)
		}

		*db.Tasks = slices.Delete(*db.Tasks, taskIndex, taskIndex+1)
	}

	return nil
}

// getTaskIndexFromNumber searches @db for existing task with given number
//
// @Returns:
//
//	task index if found of -1 if not found
func (db *InMemoryDb) getTaskIndexFromNumber(number int) int {
	t := *db.Tasks
	for i := range *db.Tasks {
		if t[i].Number == number {
			return i
		}
	}
	return -1
}
