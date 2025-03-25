package repos

import (
	"errors"
	"strings"

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

func (repo *TasksRepoStruct) Get(taskNumber int) *models.Task {
	return repo.db.Get(taskNumber)
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

func (repo *TasksRepoStruct) Add(
	name string,
	description string,
	status models.TaskStatus,
	project string,
) error {
	trimmedName := strings.TrimSpace(name)
	trimmedDescription := strings.TrimSpace(description)
	trimmedProject := strings.TrimSpace(project)

	if err := validator.TaskName(trimmedName); err != nil {
		return err
	}

	if err := validator.TaskDescription(trimmedDescription); err != nil {
		return err
	}
	return repo.db.Add(trimmedName, trimmedDescription, status, trimmedProject)
}

func (repo *TasksRepoStruct) Update(taskNumber int, task models.Task) error {
	if taskNumber < 0 {
		return InvalidTaskNumber
	}

	if err := validator.Task(task); err != nil {
		return err
	}

	return repo.db.Update(taskNumber, task)
}

func (repo *TasksRepoStruct) Delete(taskNumbers ...int) error {
	for _, number := range taskNumbers {
		if number < 0 {
			return InvalidTaskNumber
		}
	}
	return repo.db.Delete(taskNumbers...)
}
