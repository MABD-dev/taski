package db

import "github.com/mabd-dev/taski/internal/domain/models"

type Db interface {
	// Get all tasks
	GetAll() []models.Task

	// Create new task diven it's name, description and status
	// Assuming data is valid. (Should be done in domain layer)
	// @Returns:
	//     error if, actually, no case it would return nil. But it's good for later
	Add(
		name string,
		description string,
		status models.TaskStatus,
		project string,
		tags []string,
	) error

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
	// If any of the task numbers is invalid, error will be returned and nothing
	// would be deleted
	//
	// @Example:
	//   if taskNumbers = [1, 2, 3] and 2 is invalid(does not exist) 1 and 3 won't be deleted
	//
	// @Returns:
	//     error if any of task numbers is invalid
	Delete(taskNumbers ...int) error
}
