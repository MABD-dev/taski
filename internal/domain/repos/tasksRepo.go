package repos

import (
	"errors"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/validator"
)

var (
	InvalidTaskNumber = errors.New("taskNumber cannot be negative")
)

type TasksRepoStruct struct {
	db db.Db
}

var (
	TasksRepo TasksRepoStruct
)

func CreateTasksRepo(db db.Db) TasksRepoStruct {
	TasksRepo = TasksRepoStruct{
		db: db,
	}
	return TasksRepo
}

func (repo *TasksRepoStruct) List() []models.Task {
	return repo.db.List()
}

func (repo *TasksRepoStruct) Add(name string, description string, status models.TaskStatus) error {
	if err := validator.TaskName(name); err != nil {
		return err
	}

	if err := validator.TaskName(description); err != nil {
		return err
	}
	return repo.db.Add(name, description, status)
}

func (repo *TasksRepoStruct) Update(taskNumber int, name *string, description *string, status *models.TaskStatus) error {
	if taskNumber < 0 {
		return InvalidTaskNumber
	}

	if name != nil {
		if err := validator.TaskName(*name); err != nil {
			return err
		}
	}

	if description != nil {
		if err := validator.TaskName(*description); err != nil {
			return err
		}
	}

	return repo.db.Update(taskNumber, name, description, status)
}

func (repo *TasksRepoStruct) Delete(number int) error {
	if number < 0 {
		return InvalidTaskNumber
	}
	return repo.db.Delete(number)
}

func (repo *TasksRepoStruct) DeleteAll(taskNumbers []int) error {
	for _, number := range taskNumbers {
		if number < 0 {
			return InvalidTaskNumber
		}
	}
	return repo.db.DeleteAll(taskNumbers)
}
