package validator

import "github.com/mabd-dev/taski/internal/domain/models"

type MockValidator struct {
	// This value will be returned
	TaskResult       error
	NameResult       error
	DecriptionResult error
	StatusResult     error
	ProjectResult    error
}

func (v MockValidator) Task(task models.Task) error {
	return v.TaskResult
}

func (v MockValidator) TaskName(value string) error {
	return v.NameResult
}

func (v MockValidator) TaskDescription(value string) error {
	return v.DecriptionResult
}

func (v MockValidator) TaskStatus(value models.TaskStatus) error {
	return v.StatusResult
}

func (v MockValidator) TaskProject(value string) error {
	return v.ProjectResult
}
