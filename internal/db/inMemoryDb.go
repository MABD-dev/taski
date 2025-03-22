package db

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/mabd-dev/tasks/internal/models"
)

type InMemoryDb struct {
	Tasks *[]models.Task
}

func (db *InMemoryDb) List() []models.Task {
	return *db.Tasks
}

func (db *InMemoryDb) Add(name string, description string, status models.TaskStatus) {
	// WARN: do input validation

	newTaskNumber := db.findMaxTaskNumber() + 1
	newTask := models.Task{
		Number:      newTaskNumber,
		Name:        name,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	*db.Tasks = append(*db.Tasks, newTask)
}

func (db *InMemoryDb) Get(taskNumber int) *models.Task {
	for _, t := range *db.Tasks {
		if t.Number == taskNumber {
			return &t
		}
	}
	return nil
}

func (db *InMemoryDb) Update(taskNumber int, name *string, description *string, status *models.TaskStatus) error {
	// WARN: do input validation

	taskIndex := db.getTaskIndexFromNumber(taskNumber)
	if taskIndex == -1 {
		return errors.New("Could not find task")
	}

	tasks := *db.Tasks
	task := tasks[taskIndex]

	if name != nil {
		task.Name = *name
	}
	if description != nil {
		task.Description = *description
	}
	if status != nil {
		task.Status = *status
	}
	tasks[taskIndex] = task

	return nil
}

func (db *InMemoryDb) Delete(number int) error {
	taskIndex := db.getTaskIndexFromNumber(number)
	if taskIndex == -1 {
		return fmt.Errorf("Could not find task with specified number=%v", number)
	}

	*db.Tasks = slices.Delete(*db.Tasks, taskIndex, taskIndex+1)
	return nil
}

func (db *InMemoryDb) getTaskIndexFromNumber(number int) int {
	t := *db.Tasks
	for i := range *db.Tasks {
		if t[i].Number == number {
			return i
		}
	}
	return -1
}

func (db *InMemoryDb) findMaxTaskNumber() int {
	maxNumber := 1

	t := *db.Tasks
	for i := range t {
		if t[i].Number > maxNumber {
			maxNumber = t[i].Number
		}
	}
	return maxNumber
}
