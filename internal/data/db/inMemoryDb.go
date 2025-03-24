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

func (db *InMemoryDb) GetAll() []models.Task {
	return *db.Tasks
}

func (db *InMemoryDb) Add(name string, description string, status models.TaskStatus) error {

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
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	*db.Tasks = append(*db.Tasks, newTask)
	return nil
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

func (db *InMemoryDb) getTaskIndexFromNumber(number int) int {
	t := *db.Tasks
	for i := range *db.Tasks {
		if t[i].Number == number {
			return i
		}
	}
	return -1
}
