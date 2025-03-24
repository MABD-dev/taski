package repos

import (
	"errors"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/converter"
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

func (repo *TasksRepoStruct) GetAll() []models.Task {
	return repo.db.GetAll()
}

func (repo *TasksRepoStruct) ListWithFilters(statusFilters []string) []models.Task {
	tasks := repo.db.GetAll()
	if len(statusFilters) != 0 {
		statuses, err := converter.StringArrayToTaskStatus(statusFilters)
		if err != nil {
			panic(err)
		}
		tasks = converter.FilterByStatus(tasks, statuses)
	}
	return tasks
}

func (repo *TasksRepoStruct) Add(name string, description string, status models.TaskStatus) error {
	if err := validator.TaskName(name); err != nil {
		return err
	}

	if err := validator.TaskDescription(description); err != nil {
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
		if err := validator.TaskDescription(*description); err != nil {
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
