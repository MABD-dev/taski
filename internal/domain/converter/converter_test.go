package converter

import (
	"reflect"
	"testing"
	"time"

	"github.com/mabd-dev/taski/internal/domain/models"
)

func TestStringArrayToTaskStatus(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		expected  []models.TaskStatus
		expectErr bool
	}{
		{
			name:      "Valid input - all statuses",
			input:     []string{"todo", "inprogress", "done"},
			expected:  []models.TaskStatus{models.Todo, models.InProgress, models.Done},
			expectErr: false,
		},
		{
			name:      "Valid input - different cases and spaces",
			input:     []string{"  TODO  ", "InProgress", "Done "},
			expected:  []models.TaskStatus{models.Todo, models.InProgress, models.Done},
			expectErr: false,
		},
		{
			name:      "Invalid input - unknown status",
			input:     []string{"invalid"},
			expected:  []models.TaskStatus{},
			expectErr: true,
		},
		{
			name:      "Mixed valid and invalid input",
			input:     []string{"todo", "unknown", "done"},
			expected:  []models.TaskStatus{},
			expectErr: true,
		},
		{
			name:      "Empty input",
			input:     []string{},
			expected:  []models.TaskStatus{},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringArrayToTaskStatus(tt.input)

			if err == nil && tt.expectErr {
				t.Error("Expected error found nil")
			} else if err != nil && !tt.expectErr {
				t.Errorf("Expected nil found error=%v", err)
			}
			// if (err != nil) != tt.expectErr {
			// 	t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			// }

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestFilterByStatus(t *testing.T) {
	// Sample tasks
	createdAt := time.Now()
	tasks := []models.Task{
		{Number: 1, Name: "Task 1", Status: models.Todo, CreatedAt: createdAt},
		{Number: 2, Name: "Task 2", Status: models.Done, CreatedAt: createdAt},
		{Number: 3, Name: "Task 3", Status: models.Done, CreatedAt: createdAt},
		{Number: 4, Name: "Task 4", Status: models.Todo, CreatedAt: createdAt},
	}

	tests := []struct {
		name     string
		tasks    []models.Task
		statuses []models.TaskStatus
		expected []models.Task
	}{
		{
			name:     "Filter by single status",
			tasks:    tasks,
			statuses: []models.TaskStatus{models.Todo},
			expected: []models.Task{tasks[0], tasks[3]},
		},
		{
			name:     "Filter by multiple statuses",
			tasks:    tasks,
			statuses: []models.TaskStatus{models.Todo, models.Done},
			expected: []models.Task{tasks[0], tasks[1], tasks[2], tasks[3]},
		},
		{
			name:     "No matching tasks",
			tasks:    tasks,
			statuses: []models.TaskStatus{models.InProgress},
			expected: []models.Task{},
		},
		{
			name:     "Empty tasks slice",
			tasks:    []models.Task{},
			statuses: []models.TaskStatus{models.Todo},
			expected: []models.Task{},
		},
		{
			name:     "Empty statuses slice",
			tasks:    tasks,
			statuses: []models.TaskStatus{},
			expected: []models.Task{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterByStatus(tt.tasks, tt.statuses)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
