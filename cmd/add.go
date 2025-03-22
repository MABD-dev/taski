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
	RunE: func(cmd *cobra.Command, args []string) error {
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
		err = db.Add(name, description, status)
		if err != nil {
			return err
		}

		renderer.RenderTable(db.List())
		return nil
	},
}
