package db

import "github.com/mabd-dev/taski/internal/domain/models"

type Db interface {
	// Get all tasks
	GetAll() []models.Task

	// Create new task diven it's name, description and status
	// Assuming data is valid. (Should be done in domain layer)
	Add(name string, description string, status models.TaskStatus) error
	Get(taskNumber int) *models.Task
	Update(taskNumber int, name *string, description *string, status *models.TaskStatus) error

	// Delete all all with matching numbers. If any of the numbers is invalid, operation will
	// stop and return an error
	Delete(taskNumbers ...int) error
}

var (
	dbInstance Db
)

func GetDb() Db {
	if dbInstance == nil {
		dbInstance = NewPersistentDb()
	}
	return dbInstance
}
