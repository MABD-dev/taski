package models

import (
	"errors"
	"strings"
	"time"
)

type TaskStatus int

// TODO: unit test, make sure the order never changes
const (
	Todo TaskStatus = iota
	InProgress
	Done
)

// TODO: needs unit testing
// ToString, convert @TaskStatus to string
func (status TaskStatus) ToString() string {
	switch status {
	case Todo:
		return "todo"
	case InProgress:
		return "in progress"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}

// TODO: needs unit testing
// TaskStatusStrToStatus, takes a string, change it to lower case then try
// to convert it to @TaskStatus if possible, else return error
func TaskStatusStrToStatus(s string) (TaskStatus, error) {
	name := strings.ToLower(s)

	if name == "todo" {
		return Todo, nil
	}
	if name == "inprogress" {
		return InProgress, nil
	}
	if name == "done" {
		return Done, nil
	}

	err := errors.New("Could not covert from string to TaskStatus")
	return Todo, err
}

type Task struct {
	Number      int
	Name        string
	Description string
	Status      TaskStatus
	Project     string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
