package models

import (
	"testing"
)

func Test_ToString(t *testing.T) {
	if Todo.ToString() != "todo" {
		t.Fatalf("Todo.ToString(): expected 'todo' found %v", Todo.ToString())
	}

	if InProgress.ToString() != "in progress" {
		t.Fatalf("InProgress.ToString(): expected 'in progress' found %v", InProgress.ToString())
	}

	if Done.ToString() != "done" {
		t.Fatalf("Done.ToString(): expected 'done' found %v", Done.ToString())
	}

	var shit TaskStatus = 4
	if shit.ToString() != "unknown" {
		t.Fatalf("unexpected value .ToString(): expected 'unknown' found %v", shit.ToString())
	}
}

func Test_TaskStatusStrToStatus(t *testing.T) {
	tests := []struct {
		statusStr      string
		expectedStatus TaskStatus
		wantError      bool
	}{
		// todo status
		{
			statusStr:      "todo",
			expectedStatus: Todo,
			wantError:      false,
		},
		{
			statusStr:      "Todo",
			expectedStatus: Todo,
			wantError:      false,
		},
		{
			statusStr:      "ToDo",
			expectedStatus: Todo,
			wantError:      false,
		},
		{
			statusStr:      " ToDo ",
			expectedStatus: Todo,
			wantError:      false,
		},
		// in progress status
		{
			statusStr:      "inProgress",
			expectedStatus: InProgress,
			wantError:      false,
		},
		{
			statusStr:      "InProgress",
			expectedStatus: InProgress,
			wantError:      false,
		},
		{
			statusStr:      "inprogress",
			expectedStatus: InProgress,
			wantError:      false,
		},
		{
			statusStr:      "iNpRogrEss",
			expectedStatus: InProgress,
			wantError:      false,
		},
		{
			statusStr:      "INPROGRESS",
			expectedStatus: InProgress,
			wantError:      false,
		},
		{
			statusStr:      " INPROGRESS ",
			expectedStatus: InProgress,
			wantError:      false,
		},
		{
			statusStr:      "in progress",
			expectedStatus: InProgress,
			wantError:      true,
		},
		// Done status
		{
			statusStr:      "done",
			expectedStatus: Done,
			wantError:      false,
		},
		{
			statusStr:      "Done",
			expectedStatus: Done,
			wantError:      false,
		},
		{
			statusStr:      "dOne",
			expectedStatus: Done,
			wantError:      false,
		},
		{
			statusStr:      "DONE",
			expectedStatus: Done,
			wantError:      false,
		},
		// invalid status values
		{
			statusStr:      "something",
			expectedStatus: Done,
			wantError:      true,
		},
		{
			statusStr:      " ",
			expectedStatus: Done,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run("convert status str", func(t *testing.T) {
			status, err := TaskStatusStrToStatus(tt.statusStr)
			if err == nil && tt.wantError {
				t.Fatalf("(%v) Expected err found nil", tt.statusStr)
				return
			} else if err != nil && !tt.wantError {
				t.Fatalf("(%v) Expected nil found err=%v", tt.statusStr, err)
				return
			}

			if err == nil && status != tt.expectedStatus {
				t.Fatalf("Expected %v found %v for input '%v'", tt.expectedStatus, status, tt.statusStr)
			}
		})
	}
}

func Test_StatusValues(t *testing.T) {
	if Todo != 0 {
		t.Fatalf("Expected Todo to be 0 found %v", Todo)
	}
	if InProgress != 1 {
		t.Fatalf("Expected InProgress to be 1 found %v", InProgress)
	}
	if Done != 2 {
		t.Fatalf("Expected Done to be 2 found %v", Done)
	}
}
