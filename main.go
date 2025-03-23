/*
Copyright © 2025 MABD-dev <mabd.universe@gmail.com>
*/
package main

import (
	"github.com/mabd-dev/taski/cmd"
	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/domain/repos"
)

func main() {
	// Creating dependencies
	persistentDb := db.NewPersistentDb()
	repos.CreateTasksRepo(persistentDb)

	cmd.Execute()
}
