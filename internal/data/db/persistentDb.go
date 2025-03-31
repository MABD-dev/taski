package db

import (
	"github.com/mabd-dev/taski/internal/data"
	"github.com/mabd-dev/taski/internal/domain/models"
)

type PersistentDb struct {
	InMemoryDb
	storage data.LocalStorage[[]models.Task]
}

// NewPersistentDb create new @PersistentDb based on @storage and @inMemoryDb
func NewPersistentDb(
	storage *data.LocalStorage[[]models.Task],
	inMemoryDb *InMemoryDb,
) *PersistentDb {
	storage.Load(inMemoryDb.Tasks)

	return &PersistentDb{
		InMemoryDb: *inMemoryDb,
		storage:    *storage,
	}
}

// GetAll return all tasks saved in memory
// @Returns
//
//	all saved tasks in memory db
func (db *PersistentDb) GetAll() []models.Task {
	return db.InMemoryDb.GetAll()
}

// Add takes task details and add it to inMemory tasks slice.
// Input validation on name
func (db *PersistentDb) Add(
	name string,
	description string,
	status models.TaskStatus,
	project string,
	tags []string,
) error {
	err := db.InMemoryDb.Add(name, description, status, project, tags)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

// Get searches for a task with @taskNumber
//
// @Returns:
//
//	task if found or nil
func (db *PersistentDb) Get(taskNumber int) *models.Task {
	return db.InMemoryDb.Get(taskNumber)
}

// Update takes new task data and task number (ignoring that new task already has a
// task taskNumber)
// Searches for a task with @taskNumber if found it will be updated with new @task
// else return error
//
// @Returns:
//
//	update task and return nil. error if task not found based on @taskNumber
func (db *PersistentDb) Update(taskNumber int, task models.Task) error {
	err := db.InMemoryDb.Update(taskNumber, task)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

// Deelete taks list of taskNumbers and delete them.
//
// @Returns:
//
//	If any of the tasks for found  (base on it's taskNumber) will trow an error,
//	else update data and return nil
func (db *PersistentDb) Delete(taskNumbers ...int) error {
	err := db.InMemoryDb.Delete(taskNumbers...)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

// save, saves @InMemoryDb list of tasks to local file storage
func (db *PersistentDb) save() {
	db.storage.Save(*db.InMemoryDb.Tasks)
}
