package db

import "github.com/mabd-dev/taski/internal/domain/models"

type Db interface {
	// Get all tasks
	GetAll() []models.Task
	Add(name string, description string, status models.TaskStatus) error
	Get(taskNumber int) *models.Task
	Update(taskNumber int, name *string, description *string, status *models.TaskStatus) error
	Delete(number int) error
	DeleteAll(taskNumbers []int) error
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
