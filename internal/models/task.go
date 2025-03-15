package models

import "time"

type Task struct {
	Number      int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
