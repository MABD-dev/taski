package models

import (
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

type Task struct {
	Number      int
	Name        string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
