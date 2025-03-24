package db

import "github.com/mabd-dev/taski/internal/domain/models"

type Db interface {
	// Get all tasks
	GetAll() []models.Task

	// Create new task diven it's name, description and status
	// Assuming data is valid. (Should be done in domain layer)
	// @Returns:
	//     error if, actually, no case it would return nil. But it's good for later
	Add(name string, description string, status models.TaskStatus) error

	// Get task by it's number
	//
	// @Returns
	//     pointer to task, or nil if was not found
	Get(taskNumber int) *models.Task

	// Update task name, description or staus if any is provided
	//
	// @Parameters:
	//     task: take new task data and update it in db
	//
	// @Returns:
	//     error if taskNumber is invalid
	Update(taskNumber int, task models.Task) error

	// Delete all all with matching numbers.
	// Each taskNumber will be avaluated alone, when invalid number is found, function
	// will stop and return an error
	//
	// @Returns:
	//     error if any of task numbers is invalid
	Delete(taskNumbers ...int) error
}
