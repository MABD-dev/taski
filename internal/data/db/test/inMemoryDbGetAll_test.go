package db

import (
	"testing"
	"time"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/models"
)

func TestInMemoryDb_GetAll(t *testing.T) {
	tests := []struct {
		name             string
		expectedTasksLen int
		tasksToInsert    []models.Task
	}{
		{
			name:             "Empt tasks",
			expectedTasksLen: 0,
		},
		{
			name:             "Get all tasks",
			expectedTasksLen: 1,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := db.InMemoryDb{
				Tasks: &tt.tasksToInsert,
			}
			tasks := db.GetAll()
			if len(tasks) != tt.expectedTasksLen {
				t.Fatalf("Expected %v found %v tasks", tt.expectedTasksLen, len(tasks))
			}
		})
	}
}
