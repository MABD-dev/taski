/*
Copyright © 2025 MABD-dev <mabd.universe@gmail.com>
*/
package main

import (
	"github.com/mabd-dev/taski/cmd"
	"github.com/mabd-dev/taski/internal/data"
	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/repos"
)

func main() {
	storageFileName := "tasks.json"
	storage := &data.LocalStorage[[]models.Task]{FileName: storageFileName}

	tasks := make([]models.Task, 0)
	inMemoryDb := db.InMemoryDb{
		Tasks: &tasks,
	}

	persistentDb := db.NewPersistentDb(storage, &inMemoryDb)
	repos.CreateTasksRepo(persistentDb)

	cmd.Execute()
}
