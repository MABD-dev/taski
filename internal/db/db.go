package db

import (
	"fmt"
	"slices"
	"time"

	"github.com/mabd-dev/tasks/internal/models"
)

type Db struct {
	Tasks *[]models.Task
}

var (
	dbInstance *Db
)

func GetDb() *Db {
	tasks := make([]models.Task, 0)
	dbInstance = &Db{
		Tasks: &tasks,
	}
	return dbInstance
}

func (db *Db) List() {
	t := *db.Tasks

	fmt.Println("\n****************")
	for i := range *db.Tasks {
		fmt.Printf("Task: number=%v, name=`%v`, description=`%v`, createdAt=%v", t[i].Number,
			t[i].Name, t[i].Description, t[i].CreatedAt.Format(time.RFC1123))
	}
	fmt.Println("\n****************")
}

func (db *Db) Add(name string, description string) {
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

func (db *Db) Delete(number int) error {
	taskIndex := db.getTaskIndexFromNumber(number)
	if taskIndex == -1 {
		return fmt.Errorf("Could not find task with specified number=%v", number)
	}

	*db.Tasks = slices.Delete(*db.Tasks, taskIndex, taskIndex+1)
	return nil
}

func (db *Db) getTaskIndexFromNumber(number int) int {
	t := *db.Tasks
	for i := range *db.Tasks {
		if t[i].Number == number {
			return i
		}
	}
	return -1
}
