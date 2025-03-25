package cmd

import (
	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
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

		project, err := cmd.Flags().GetString("project")
		if err != nil {
			panic(err)
		}

		name := args[0]

		err = repos.TasksRepo.Add(name, description, status, project)
		if err != nil {
			return err
		}

		ui.RenderKanbanBoard(repos.TasksRepo.GetAll())
		return nil
	},
}
