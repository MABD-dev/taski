package cmd

import (
	"github.com/mabd-dev/tasks/internal/db"
	"github.com/mabd-dev/tasks/internal/models"
	"github.com/mabd-dev/tasks/internal/renderer"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [task name] [task description]",
	Short: "Add new task",
	Long:  "Add new task to the list with default completion value to false",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		statusStr, err := cmd.Flags().GetString("status")
		if err != nil {
			panic(err)
		}
		status, err := models.TaskStatusStrToStatus(statusStr)
		if err != nil {
			panic(err)
		}

		description, err := cmd.Flags().GetString("description")
		if err != nil {
			panic(err)
		}

		name := args[0]

		db := db.GetDb()
		db.Add(name, description, status)

		renderer.RenderTable(db.List())
	},
}
