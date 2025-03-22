package models

import (
	"errors"
	"strings"
	"time"
)

type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

func (status TaskStatus) ToString() string {
	switch status {
	case Todo:
		return "todo"
	case InProgress:
		return "inProgress"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}

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
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
