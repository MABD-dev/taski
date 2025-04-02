package db

import (
	"testing"
	"time"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/models"
)

func TestInMemoryDb_Delete(t *testing.T) {
	tests := []struct {
		name                string
		taskNumbersToDelete []int
		wantErr             bool
		tasksToInsert       []models.Task
	}{
		{
			name:                "delete nothing",
			taskNumbersToDelete: []int{},
			wantErr:             false,
		},
		{
			name:                "Delete unexisting task",
			taskNumbersToDelete: []int{1},
			wantErr:             true,
		},
		{
			name:                "Delete 1 task",
			taskNumbersToDelete: []int{1},
			wantErr:             false,
			tasksToInsert: []models.Task{
				{
					Number:      1,
					Name:        "a",
					Description: "d",
					Status:      models.Todo,
					CreatedAt:   time.Now(),
					UpdatedAt:   nil,
				},
			},
		},
		{
			name:                "Delete multiple tasks",
			taskNumbersToDelete: []int{1, 2},
			wantErr:             false,
			tasksToInsert: []models.Task{
				{
					Number:      1,
					Name:        "a",
					Description: "d",
					Status:      models.Todo,
					CreatedAt:   time.Now(),
					UpdatedAt:   nil,
				},
				{
					Number:      2,
					Name:        "a",
					Description: "d",
					Status:      models.Todo,
					CreatedAt:   time.Now(),
					UpdatedAt:   nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := db.InMemoryDb{
				Tasks: &tt.tasksToInsert,
			}
			gotErr := db.Delete(tt.taskNumbersToDelete...)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Delete() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Delete() succeeded unexpectedly")
			}

			allTasks := db.GetAll()
			if len(allTasks) != 0 {
				t.Fatalf("Expected number of task 0 found %v", len(allTasks))
			}
		})
	}
}

func TestDeleteMultipleTasksWithSomeInvalid(t *testing.T) {
	tasks := []models.Task{
		{
			Number:      1,
			Name:        "a",
			Description: "d",
			Status:      models.Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		},
		{
			Number:      2,
			Name:        "a",
			Description: "d",
			Status:      models.Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		},
	}
	oldTasksLen := len(tasks)

	db := db.InMemoryDb{
		Tasks: &tasks,
	}
	err := db.Delete(1, 3, 2)
	if err == nil {
		t.Fatal("Expected error found nil")
	}

	tasksInDb := db.GetAll()
	if len(tasksInDb) != oldTasksLen {
		t.Fatalf("Expected %v tasks in db found %v", oldTasksLen, len(tasksInDb))
	}
}
