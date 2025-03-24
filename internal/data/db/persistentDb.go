package db

import (
	"github.com/mabd-dev/taski/internal/data"
	"github.com/mabd-dev/taski/internal/domain/models"
)

type PersistentDb struct {
	InMemoryDb
	storage data.LocalStorage[[]models.Task]
}

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

func (db *PersistentDb) GetAll() []models.Task {
	return db.InMemoryDb.GetAll()
}

func (db *PersistentDb) Add(name string, description string, status models.TaskStatus) error {
	err := db.InMemoryDb.Add(name, description, status)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

func (db *PersistentDb) Get(taskNumber int) *models.Task {
	return db.InMemoryDb.Get(taskNumber)
}

func (db *PersistentDb) Update(taskNumber int, task models.Task) error {
	err := db.InMemoryDb.Update(taskNumber, task)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

func (db *PersistentDb) Delete(taskNumbers ...int) error {
	err := db.InMemoryDb.Delete(taskNumbers...)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

func (db *PersistentDb) save() {
	db.storage.Save(*db.InMemoryDb.Tasks)
}
