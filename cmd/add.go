package cmd

import (
	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/presentation"
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

		err = repos.TasksRepo.Add(name, description, status)
		if err != nil {
			return err
		}

		presentation.RenderTable(repos.TasksRepo.List())
		return nil
	},
}
