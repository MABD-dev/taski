package db

import "github.com/mabd-dev/tasks/internal/models"

type Db interface {
	List() []models.Task
	Add(name string, description string)
	Delete(number int) error
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
