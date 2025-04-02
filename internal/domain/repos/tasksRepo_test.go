package repos

import (
	"errors"
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
// *********** Add Task Testing *************
// ******************************************
// ******************************************

func TestAddTask_ValidationIsWorking(t *testing.T) {
	taskToInsert := models.Task{
		Number:      1,
		Name:        "Something",
		Description: "some description",
		Status:      models.Done,
		Project:     "Something",
		Tags:        []string{},
		CreatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		tasksToInsert models.Task
		validator     validator.Validator
		wantErr       bool
	}{
		{
			name:          "Test: Task name validator is working",
			tasksToInsert: taskToInsert,
			validator: validator.MockValidator{
				NameResult: errors.New("testing name validator"),
			},
			wantErr: true,
		},
		{
			name:          "Test: Task description validator is working",
			tasksToInsert: taskToInsert,
			validator: validator.MockValidator{
				DecriptionResult: errors.New("testing description validator"),
			},
			wantErr: true,
		},
		{
			name:          "Test: Task status validator is working",
			tasksToInsert: taskToInsert,
			validator: validator.MockValidator{
				StatusResult: errors.New("testing status validator"),
			},
			wantErr: true,
		},
		{
			name:          "Test: Task project validator is working",
			tasksToInsert: taskToInsert,
			validator: validator.MockValidator{
				ProjectResult: errors.New("testing project validator"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tasks := []models.Task{}
			db := db.InMemoryDb{
				Tasks: &tasks,
			}
			repo := CreateTasksRepo(&db, tt.validator)

			task := tt.tasksToInsert
			err := repo.Add(task.Name, task.Description, task.Status, task.Project, task.Tags)

			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}

			fetchedTasks := repo.GetAll()
			expectedTasksLen := 1
			if tt.wantErr {
				expectedTasksLen = 0
			}
			if len(fetchedTasks) != expectedTasksLen {
				t.Fatalf("Expected %v task in db found %v", expectedTasksLen, len(fetchedTasks))
			}
		})
	}
}

func TestAddTask_NameWhiteSpaceIsRemoved(t *testing.T) {
	tests := []struct {
		name         string
		taskToInsert models.Task
		expectedTask models.Task
		wantErr      bool
	}{
		{
			name: "Task name whitespace is removed",
			taskToInsert: models.Task{
				Number: 1,
				Name:   " a ",
			},
			expectedTask: models.Task{
				Number: 1,
				Name:   "a",
			},
		},
		{
			name: "Task description whitespace is removed",
			taskToInsert: models.Task{
				Number:      1,
				Name:        "a",
				Description: " b ",
			},
			expectedTask: models.Task{
				Number:      1,
				Name:        "a",
				Description: "b",
			},
		},
		{
			name: "Task project whitespace is removed",
			taskToInsert: models.Task{
				Number:  1,
				Name:    "a",
				Project: " b ",
			},
			expectedTask: models.Task{
				Number:  1,
				Name:    "a",
				Project: "b",
			},
		},
		{
			name: "Task tags whitespace is removed",
			taskToInsert: models.Task{
				Number: 1,
				Name:   "a",
				Tags:   []string{" a", "b ", "   c   "},
			},
			expectedTask: models.Task{
				Number: 1,
				Name:   "a",
				Tags:   []string{"a", "b", "c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := createRepoNoData()

			task := tt.taskToInsert
			err := repo.Add(task.Name, task.Description, task.Status, task.Project, task.Tags)

			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}

			allTasks := repo.GetAll()
			expectedTasksLen := 1
			if tt.wantErr {
				expectedTasksLen = 0
			}
			if len(allTasks) != expectedTasksLen {
				t.Fatalf("Expected %v task in db found %v", expectedTasksLen, len(allTasks))
			}

			fetchedTask := allTasks[0]
			if fetchedTask.Number != tt.expectedTask.Number {
				t.Errorf("Expected task number %v found %v", tt.expectedTask.Number, fetchedTask.Number)
			}

			if fetchedTask.Name != tt.expectedTask.Name {
				t.Errorf("Expected task name %v found %v", tt.expectedTask.Name, fetchedTask.Name)
			}

			if fetchedTask.Description != tt.expectedTask.Description {
				t.Errorf("Expected task description %v found %v", tt.expectedTask.Description, fetchedTask.Description)
			}

			if fetchedTask.Status != tt.expectedTask.Status {
				t.Errorf("Expected task status %v found %v", tt.expectedTask.Status, fetchedTask.Status)
			}

			if fetchedTask.Project != tt.expectedTask.Project {
				t.Errorf("Expected task project %v found %v", tt.expectedTask.Project, fetchedTask.Project)
			}

			if len(fetchedTask.Tags) != len(tt.expectedTask.Tags) {
				t.Errorf("Expected task tags size %v found %v", len(tt.expectedTask.Tags), (fetchedTask.Tags))
			}

			for i := range fetchedTask.Tags {
				if fetchedTask.Tags[i] != tt.expectedTask.Tags[i] {
					t.Errorf("Expected task tag[%v] %v found %v", i, len(tt.expectedTask.Tags), (fetchedTask.Tags))
				}
			}
		})
	}
}

// ******************************************
// ******************************************
// ********** Update Task Testing ***********
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

func TestUpdate_ValidatorIsWorking(t *testing.T) {
	validatorError := errors.New("task validator is working")
	validator := validator.MockValidator{
		TaskResult: validatorError,
	}
	tasks := []models.Task{
		{
			Number: 1,
			Name:   "Something",
		},
	}
	db := db.InMemoryDb{
		Tasks: &tasks,
	}
	repo := CreateTasksRepo(&db, validator)

	err := repo.Update(1, models.Task{
		Number: 1,
		Name:   "Something else",
	})
	if err != validatorError {
		t.Fatalf("Expected error=%v, found=%v", validatorError, err)
	}
}

// TODO: to be added later when InMemoryDb is updated

func TestUpdateTaskNumberHasNoEffect(t *testing.T) {
	tasks := []models.Task{
		{
			Number: 1,
			Name:   "Something",
		},
	}
	repo := createRepo(tasks)

	err := repo.Update(1, models.Task{
		Number: 2,
		Name:   "Something",
	})
	if err != nil {
		t.Fatalf("Expected no error found error=%v", err)
		return
	}

	fetchedTask := repo.Get(1)
	if fetchedTask == nil {
		t.Fatal("Expected task found nil")
	}
}

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
