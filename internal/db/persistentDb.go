package db

import (
	"github.com/mabd-dev/tasks/internal/models"
)

type PersistentDb struct {
	InMemoryDb
	storage LocalStorage[[]models.Task]
}

func NewPersistentDb() *PersistentDb {
	storage := NewLocalStorage[[]models.Task]("data/tasks.json")
	tasks := make([]models.Task, 0)
	inMemoryDb := &InMemoryDb{
		Tasks: &tasks,
	}
	storage.Load(inMemoryDb.Tasks)

	return &PersistentDb{
		InMemoryDb: *inMemoryDb,
		storage:    *storage,
	}
}

func (db *PersistentDb) List() []models.Task {
	return db.InMemoryDb.List()
}

func (db *PersistentDb) Add(name string, description string) {
	db.InMemoryDb.Add(name, description)
	db.storage.Save(*db.InMemoryDb.Tasks)
}

func (db *PersistentDb) Delete(number int) error {
	err := db.InMemoryDb.Delete(number)
	if err != nil {
		return err
	}
	db.storage.Save(*db.InMemoryDb.Tasks)
	return nil
}
