package validator

import (
	"errors"
	"unicode/utf8"

	"github.com/mabd-dev/taski/internal/domain/models"
)

func TaskName(value string) error {
	nameLen := utf8.RuneCountInString(value)
	if nameLen == 0 {
		return errors.New("name cannot be empty")
	}
	if nameLen > 50 {
		return errors.New("name must be less than 50 characters")
	}
	return nil
}

func TaskDescription(value string) error {
	descriptionLen := utf8.RuneCountInString(value)
	if descriptionLen > 200 {
		return errors.New("description must be less than 200 characters")
	}

	return nil
}

func TaskStatus(value models.TaskStatus) error {
	if value != models.Todo && value != models.InProgress && value != models.Done {
		return errors.New("invalid status")
	}
	return nil
}
