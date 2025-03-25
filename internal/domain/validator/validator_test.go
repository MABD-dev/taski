package validator

import (
	"strings"
	"testing"

	"github.com/mabd-dev/taski/internal/domain/config"
	"github.com/mabd-dev/taski/internal/domain/models"
)

func TestValidator_Name(t *testing.T) {
	tests := []struct {
		name     string
		taskName string
		wantErr  bool
	}{
		{
			name:     "Blank task name",
			taskName: "",
			wantErr:  true,
		},
		{
			name:     "Whitespace task name",
			taskName: " ",
			wantErr:  true,
		},
		{
			name:     "Very long task name",
			taskName: strings.Repeat("A", config.TaskNameMaxLen+1),
			wantErr:  true,
		},
		{
			name:     "Valid task name",
			taskName: "asdf",
			wantErr:  false,
		},
		{
			name:     "Valid long task name",
			taskName: strings.Repeat("A", config.TaskNameMaxLen),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := TaskName(tt.taskName)
			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}
		})
	}

}

func TestValidator_Description(t *testing.T) {
	tests := []struct {
		name        string
		description string
		wantErr     bool
	}{
		{
			name:        "Blank description",
			description: "",
			wantErr:     false,
		},
		{
			name:        "Whitespace description",
			description: " ",
			wantErr:     false,
		},
		{
			name:        "Very long description",
			description: strings.Repeat("A", config.TaskDescriptionMaxLen+1),
			wantErr:     true,
		},
		{
			name:        "Valid description",
			description: "asdf",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := TaskDescription(tt.description)
			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}
		})
	}

}

func TestValidator_TaskStatus(t *testing.T) {
	tests := []struct {
		name       string
		taskStatus models.TaskStatus
		wantErr    bool
	}{
		{
			name:       "todo status",
			taskStatus: models.Todo,
			wantErr:    false,
		},
		{
			name:       "in progress status",
			taskStatus: models.InProgress,
			wantErr:    false,
		},
		{
			name:       "done status",
			taskStatus: models.Done,
			wantErr:    false,
		},
		{
			name:       "invalid status",
			taskStatus: 3,
			wantErr:    true,
		},
		{
			name:       "invalid status",
			taskStatus: -1,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := TaskStatus(tt.taskStatus)
			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}
		})
	}

}

func TestValidator_Project(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		wantErr     bool
	}{
		{
			name:        "Blank project name",
			projectName: "",
			wantErr:     false,
		},
		{
			name:        "Whitespace project name",
			projectName: " ",
			wantErr:     false,
		},
		{
			name:        "Very long project name",
			projectName: strings.Repeat("A", config.TaskProjectMaxLen+1),
			wantErr:     true,
		},
		{
			name:        "Valid long project name",
			projectName: strings.Repeat("A", config.TaskProjectMaxLen),
			wantErr:     false,
		},
		{
			name:        "Valid project name",
			projectName: "asdf",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := TaskProject(tt.projectName)
			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}
		})
	}

}

func TestValidator_Task(t *testing.T) {
	tests := []struct {
		name    string
		task    models.Task
		wantErr bool
	}{
		{
			name: "Invalid task number",
			task: models.Task{
				Number: -1,
			},
			wantErr: true,
		},
		{
			name: "Invalid task number",
			task: models.Task{
				Number: 0,
			},
			wantErr: true,
		},
		{
			name: "Invalid task name",
			task: models.Task{
				Number: 1,
				Name:   "",
			},
			wantErr: true,
		},
		{
			name: "Invalid task name",
			task: models.Task{
				Number: 1,
				Name:   " ",
			},
			wantErr: true,
		},
		{
			name: "Invalid task name",
			task: models.Task{
				Number: 1,
				Name:   strings.Repeat("A", config.TaskNameMaxLen+1),
			},
			wantErr: true,
		},
		{
			name: "Invalid task description",
			task: models.Task{
				Number:      1,
				Name:        "A",
				Description: strings.Repeat("A", config.TaskDescriptionMaxLen+1),
			},
			wantErr: true,
		},
		{
			name: "Invalid task status",
			task: models.Task{
				Number:      1,
				Name:        "A",
				Description: "",
				Status:      3,
			},
			wantErr: true,
		},
		{
			name: "Invalid task project",
			task: models.Task{
				Number:      1,
				Name:        "A",
				Description: "",
				Status:      2,
				Project:     strings.Repeat("A", config.TaskProjectMaxLen+1),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Task(tt.task)
			if err == nil && tt.wantErr {
				t.Fatal("Expected error found nil")
			} else if err != nil && !tt.wantErr {
				t.Fatalf("Expected nil found error=%v", err)
			}
		})
	}

}
