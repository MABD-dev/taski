package repos

import (
	"strings"
	"testing"
	"time"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/config"
	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/validator"
)

func createRepoNoData() TasksRepoStruct {
	tasks := make([]models.Task, 0)
	validator := validator.ValidatorImpl{}
	db := db.InMemoryDb{
		Tasks: &tasks,
	}
	return CreateTasksRepo(&db, validator)
}

func createRepo(tasks []models.Task) TasksRepoStruct {
	validator := validator.ValidatorImpl{}
	db := db.InMemoryDb{
		Tasks: &tasks,
	}
	return CreateTasksRepo(&db, validator)
}

// ******************************************
// ******************************************
// ************ Get All Tests ***************
// ******************************************
// ******************************************

func TestGetAllTasks_emptyDB(t *testing.T) {
	repo := createRepoNoData()

	tasks := repo.GetAll()
	if len(tasks) != 0 {
		t.Fatalf("Expected 0 tasks found %v", len(tasks))
	}
}

func TestGetAllTasks_With(t *testing.T) {
	tasks := []models.Task{
		{
			Number:      1,
			Name:        "Something",
			Description: "some description",
			Status:      models.Done,
			Project:     "Something",
			Tags:        []string{},
			CreatedAt:   time.Now(),
		},
	}
	repo := createRepo(tasks)

	fetchedTasks := repo.GetAll()
	if len(fetchedTasks) != len(tasks) {
		t.Fatalf("Expected %v tasks found %v", len(tasks), len(fetchedTasks))
	}
}

// ******************************************
// ******************************************
// ************ Get Task Tests **************
// ******************************************
// ******************************************

func TestGetGetTask_emptyDB(t *testing.T) {
	repo := createRepoNoData()

	task := repo.Get(0)
	if task != nil {
		t.Fatalf("Expected nil found task=%v", task)
	}
}

func TestGetTasks_With(t *testing.T) {
	originalTask := models.Task{
		Number:      1,
		Name:        "Something",
		Description: "some description",
		Status:      models.Done,
		Project:     "Something",
		Tags:        []string{},
		CreatedAt:   time.Now(),
	}
	tasks := []models.Task{originalTask}
	repo := createRepo(tasks)

	task := repo.Get(1)
	if task == nil {
		t.Fatalf("Expected task=%v found nil", tasks)
	}
	if task.Name != originalTask.Name {
		t.Fatalf("Expected task name=%v found %v", originalTask.Name, task.Name)
	}
	if task.Description != originalTask.Description {
		t.Fatalf("Expected task description=%v found %v", originalTask.Description, task.Description)
	}
	if task.Status != originalTask.Status {
		t.Fatalf("Expected task status=%v found %v", originalTask.Status, task.Status)
	}
}

// ******************************************
// ******************************************
// ************ Update Testing **************
// ******************************************
// ******************************************

func TestUpdateUnexistingTask(t *testing.T) {
	task := models.Task{
		Number:      1,
		Name:        "Something",
		Description: "some description",
		Status:      models.Done,
		Project:     "Something",
		Tags:        []string{},
		CreatedAt:   time.Now(),
	}
	repo := createRepoNoData()
	err := repo.Update(1, task)
	if err == nil {
		t.Fatal("Expected error found nil")
	}
}

// TODO: to be added later when InMemoryDb is updated
//
// func TestUpdateTaskNumberHasNoEffect(t *testing.T) {
// 	task := models.Task{
// 		Number:      1,
// 		Name:        "Something",
// 		Description: "some description",
// 		Status:      models.Done,
// 		Project:     "Something",
// 		Tags:        []string{},
// 		CreatedAt:   time.Now(),
// 	}
// 	tasks := []models.Task{task}
// 	repo := createRepo(tasks)
//
// 	err := repo.Update(1, task)
// 	if err != nil {
// 		t.Fatalf("Expected no error found error=%v", err)
// 		return
// 	}
//
// 	fetchedTask := repo.Get(1)
// 	if fetchedTask == nil {
// 		t.Fatal("Expected task found nil")
// 	}
// }

func TestUpdateValidationIsWorking(t *testing.T) {
	existingTask := models.Task{
		Number:      1,
		Name:        "Something",
		Description: "Something Else",
		Status:      models.Done,
		Project:     "A Project",
		Tags:        []string{},
	}
	tests := []struct {
		name    string
		wantErr bool
		task    models.Task
	}{
		{
			name:    "Invalid Task number 0",
			wantErr: true,
			task: models.Task{
				Number: 0,
			},
		},
		{
			name:    "Invalid Task number -1",
			wantErr: true,
			task: models.Task{
				Number: -1,
			},
		},
		{
			name:    "Invalid task name: empty",
			wantErr: true,
			task: models.Task{
				Number: 1,
				Name:   "",
			},
		},
		{
			name:    "Invalid task name: only whitespace",
			wantErr: true,
			task: models.Task{
				Number: 1,
				Name:   " ",
			},
		},
		{
			name:    "Invalid task name: very long",
			wantErr: true,
			task: models.Task{
				Number: 1,
				Name:   strings.Repeat("A", config.TaskNameMaxLen+1),
			},
		},
		{
			name:    "Invalid task description: very long",
			wantErr: true,
			task: models.Task{
				Number:      1,
				Name:        "A",
				Description: strings.Repeat("A", config.TaskDescriptionMaxLen+1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks := []models.Task{existingTask}
			db := createRepo(tasks)
			err := db.Update(tt.task.Number, tt.task)
			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}

			fetchedTask := db.Get(1)
			if fetchedTask == nil {
				t.Fatal("Expected task found il")
			}

			err = validator.ValidatorImpl{}.Task(*fetchedTask)
			if err != nil {
				t.Fatalf("Expected task to be valid found otherwise, task=%v, err=%v", fetchedTask, err)
			}
		})
	}
}

// ******************************************
// ******************************************
// *********** Delete Task Tests ************
// ******************************************
// ******************************************

func TestDeleteUnexistingTask(t *testing.T) {
	repo := createRepoNoData()

	err := repo.Delete(0)
	if err == nil {
		t.Fatal("Expected error found nil")
	}

	tasks := repo.GetAll()
	if len(tasks) > 0 {
		t.Fatalf("Expected 0 tasks found %v", len(tasks))
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []models.Task{
		{
			Number:      1,
			Name:        "Something",
			Description: "some description",
			Status:      models.Done,
			Project:     "Something",
			Tags:        []string{},
			CreatedAt:   time.Now(),
		},
	}
	repo := createRepo(tasks)

	err := repo.Delete(1)
	if err != nil {
		t.Fatalf("Expected nill error found error=%v", err)
	}

	tasks = repo.GetAll()
	if len(tasks) > 0 {
		t.Fatalf("Expected 0 tasks found %v", len(tasks))
	}
}

func TestDeletetingMultipleTasks(t *testing.T) {
	tasks := []models.Task{
		{
			Number:      1,
			Name:        "Something",
			Description: "some description",
			Status:      models.Done,
			Project:     "Something",
			Tags:        []string{},
			CreatedAt:   time.Now(),
		},
		{
			Number:      2,
			Name:        "Something",
			Description: "some description",
			Status:      models.Done,
			Project:     "Something",
			Tags:        []string{},
			CreatedAt:   time.Now(),
		},
	}
	repo := createRepo(tasks)

	err := repo.Delete(1, 2)
	if err != nil {
		t.Fatalf("Expected nill error found error=%v", err)
	}

	tasks = repo.GetAll()
	if len(tasks) > 0 {
		t.Fatalf("Expected 0 tasks found %v", len(tasks))
	}
}

func TestDeletetingMultipleTasksWithUnexistingOne(t *testing.T) {
	tasks := []models.Task{
		{
			Number:      1,
			Name:        "Something",
			Description: "some description",
			Status:      models.Done,
			Project:     "Something",
			Tags:        []string{},
			CreatedAt:   time.Now(),
		},
		{
			Number:      2,
			Name:        "Something",
			Description: "some description",
			Status:      models.Done,
			Project:     "Something",
			Tags:        []string{},
			CreatedAt:   time.Now(),
		},
	}
	repo := createRepo(tasks)

	err := repo.Delete(1, 3, 2)
	if err == nil {
		t.Fatal("Expected err error nil")
	}

	// when 3 fails, 1 should not be deleted
	tasks = repo.GetAll()
	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks found %v", len(tasks))
	}
}
