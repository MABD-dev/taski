package db

import (
	"github.com/mabd-dev/tasks/internal/models"
	"github.com/mabd-dev/tasks/internal/services"
)

type PersistentDb struct {
	InMemoryDb
	storage services.LocalStorage[[]models.Task]
}

func NewPersistentDb() *PersistentDb {
	storage := services.NewLocalStorage[[]models.Task]("data/tasks.json")
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

func (db *PersistentDb) Update(taskNumber int, name *string, description *string, status *models.TaskStatus) error {
	err := db.InMemoryDb.Update(taskNumber, name, description, status)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

func (db *PersistentDb) Delete(number int) error {
	err := db.InMemoryDb.Delete(number)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

func (db *PersistentDb) DeleteAll(taskNumbers []int) error {
	err := db.InMemoryDb.DeleteAll(taskNumbers)
	if err != nil {
		return err
	}
	db.save()
	return nil
}

func (db *PersistentDb) save() {
	db.storage.Save(*db.InMemoryDb.Tasks)
}
