package db

import (
	"testing"
	"time"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/models"
)

type taskToUpdate struct {
	name        *string
	description *string
	status      *models.TaskStatus
}

func TestInMemoryDb_Update(t *testing.T) {
	newName := "aa"
	newDescription := "dd"
	newStatus := models.Done

	defaultTask := models.Task{
		Number:      1,
		Name:        "a",
		Description: "d",
		Status:      models.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	tests := []struct {
		name         string
		taskToInsert models.Task
		expectedTask models.Task
		taskNumber   int
		wantErr      bool
	}{
		{
			name:       "Update unexisting task",
			taskNumber: 1,
			wantErr:    true,
		},
		{
			name:         "Update nothing",
			taskToInsert: defaultTask,
			expectedTask: defaultTask,
			taskNumber:   1,
			wantErr:      false,
		},
		{
			name:         "Update name only",
			taskToInsert: defaultTask,
			expectedTask: models.Task{
				Number:      1,
				Name:        newName,
				Description: "d",
				Status:      models.Todo,
				CreatedAt:   time.Now(),
				UpdatedAt:   nil,
			},
			taskNumber: 1,
			wantErr:    false,
		},
		{
			name:         "Update description only",
			taskToInsert: defaultTask,
			expectedTask: models.Task{
				Number:      1,
				Name:        "a",
				Description: newDescription,
				Status:      models.Todo,
				CreatedAt:   time.Now(),
				UpdatedAt:   nil,
			},
			taskNumber: 1,
			wantErr:    false,
		},
		{
			name:         "Update status only",
			taskToInsert: defaultTask,
			expectedTask: models.Task{
				Number:      1,
				Name:        "a",
				Description: "d",
				Status:      newStatus,
				CreatedAt:   time.Now(),
				UpdatedAt:   nil,
			},
			taskNumber: 1,
			wantErr:    false,
		},
		{
			name:         "Update tas",
			taskToInsert: defaultTask,
			expectedTask: models.Task{
				Number:      1,
				Name:        newName,
				Description: newDescription,
				Status:      newStatus,
				CreatedAt:   time.Now(),
				UpdatedAt:   nil,
			},
			taskNumber: 1,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := db.InMemoryDb{
				Tasks: &[]models.Task{tt.taskToInsert},
			}
			err := db.Update(tt.taskNumber, tt.expectedTask)
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("Expected error=%v", err)
				}
			}

			fetchedTask := db.GetAll()[0]
			if fetchedTask.Number != tt.expectedTask.Number {
				t.Errorf("Expected task Number %v found %v", tt.expectedTask.Number, fetchedTask.Number)
			}

			if fetchedTask.Name != tt.expectedTask.Name {
				t.Errorf("Expected task Name %v found %v", tt.expectedTask.Name, fetchedTask.Name)
			}

			if fetchedTask.Description != tt.expectedTask.Description {
				t.Errorf("Expected task Description %v found %v", tt.expectedTask.Description, fetchedTask.Description)
			}

			if fetchedTask.Status != tt.expectedTask.Status {
				t.Errorf("Expected task Status %v found %v", tt.expectedTask.Status, fetchedTask.Status)
			}
		})
	}
}
