package db

import (
	"testing"
	"time"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/models"
)

func TestInMemoryDb_Get(t *testing.T) {
	tests := []struct {
		name         string
		taskToInsert models.Task
		expectNil    bool

		// to fetch by
		taskNumber int
	}{
		{
			name:       "Fetching unexisting task",
			expectNil:  true,
			taskNumber: 1,
		},
		{
			name: "Get task",
			taskToInsert: models.Task{
				Number:      1,
				Name:        "a",
				Description: "d",
				Status:      models.Todo,
				CreatedAt:   time.Now(),
				UpdatedAt:   nil,
			},
			taskNumber: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := db.InMemoryDb{
				Tasks: &[]models.Task{tt.taskToInsert},
			}

			task := db.Get(tt.taskNumber)
			if task == nil {
				if !tt.expectNil {
					t.Fatalf("Expected task found nil")
				}
				return
			}

			if task.Number != 1 {
				t.Errorf("Expected task number 1 found %v", task.Number)
			}

			if task.Name != tt.taskToInsert.Name {
				t.Errorf("Expected task name %v found %v", tt.taskToInsert.Name, task.Name)
			}

			if task.Description != tt.taskToInsert.Description {
				t.Errorf("Expected task description %v found %v", tt.taskToInsert.Description, task.Description)
			}

			if task.Status != task.Status {
				t.Errorf("Expected task Status %v found %v", task.Status, task.Status)
			}
		})
	}
}
