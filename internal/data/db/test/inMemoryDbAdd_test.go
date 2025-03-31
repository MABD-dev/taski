package db

import (
	"testing"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/models"
)

type taskToInsertData struct {
	name        string
	description string
	status      models.TaskStatus
	project     string
	tags        []string
}

func TestInMemoryDb_Add(t *testing.T) {
	tests := []struct {
		name             string
		expectedTasksLen int
		taskToInsert     taskToInsertData
		wantErr          bool
	}{
		{
			name:             "Add one task",
			expectedTasksLen: 1,
			taskToInsert: taskToInsertData{
				name:        "a",
				description: "description",
				status:      models.Todo,
			},
		},
		{
			name:             "Add one task with project",
			expectedTasksLen: 1,
			taskToInsert: taskToInsertData{
				name:        "a",
				description: "description",
				status:      models.Todo,
				project:     "taski",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := db.InMemoryDb{
				Tasks: &[]models.Task{},
			}
			err := db.Add(
				tt.taskToInsert.name,
				tt.taskToInsert.description,
				tt.taskToInsert.status,
				tt.taskToInsert.project,
				tt.taskToInsert.tags,
			)
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("Expected error=%v", err)
				}
			}

			allTasks := db.GetAll()
			if len(allTasks) != tt.expectedTasksLen {
				t.Fatalf("Expected %v found %v tasks", tt.expectedTasksLen, len(allTasks))
			}

			fetchedTask := allTasks[0]

			if fetchedTask.Number != 1 {
				t.Errorf("Expected task number 1 found %v", fetchedTask.Number)
			}

			if fetchedTask.Name != tt.taskToInsert.name {
				t.Errorf("Expected task name %v found %v", tt.taskToInsert.name, fetchedTask.Name)
			}

			if fetchedTask.Description != tt.taskToInsert.description {
				t.Errorf("Expected task description %v found %v", tt.taskToInsert.description, fetchedTask.Description)
			}

			if fetchedTask.Status != tt.taskToInsert.status {
				t.Errorf("Expected task status %v found %v", tt.taskToInsert.status, fetchedTask.Status)
			}

			if fetchedTask.Project != tt.taskToInsert.project {
				t.Errorf("Expected task project %v found %v", tt.taskToInsert.project, fetchedTask.Project)
			}
		})
	}
}
