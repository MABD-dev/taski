package db

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mabd-dev/tasks/internal/models"
)

type InMemoryDb struct {
	Tasks *[]models.Task
}

func (db *InMemoryDb) List() {
	t := *db.Tasks

	table := table.New(os.Stdout)
	table.SetHeaders("#", "Name", "Description", "Created At")

	for i := range *db.Tasks {
		task := t[i]
		table.AddRow(strconv.Itoa(task.Number), task.Name, task.Description, task.CreatedAt.Format(time.RFC1123))
	}
	table.Render()
}

func (db *InMemoryDb) Add(name string, description string) {
	// WARN: do input validation on name and description

	nOfTasks := len(*db.Tasks)
	newTask := models.Task{
		Number:      nOfTasks + 1,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	*db.Tasks = append(*db.Tasks, newTask)
}

func (db *InMemoryDb) Delete(number int) error {
	taskIndex := db.getTaskIndexFromNumber(number)
	if taskIndex == -1 {
		return fmt.Errorf("Could not find task with specified number=%v", number)
	}

	*db.Tasks = slices.Delete(*db.Tasks, taskIndex, taskIndex+1)
	return nil
}

func (db *InMemoryDb) getTaskIndexFromNumber(number int) int {
	t := *db.Tasks
	for i := range *db.Tasks {
		if t[i].Number == number {
			return i
		}
	}
	return -1
}
