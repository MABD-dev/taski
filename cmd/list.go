package cmd

import (
	"github.com/mabd-dev/taski/internal/domain/converter"
	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all your tasks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		statuses, err := cmd.Flags().GetStringArray("status")
		if err != nil {
			panic(err)
		}

		tasks := repos.TasksRepo.List()

		if len(statuses) != 0 {
			statuses, err := converter.StringArrayToTaskStatus(statuses)
			if err != nil {
				panic(err)
			}
			tasks = converter.FilterByStatus(tasks, statuses)
		}

		ui.RenderTable(tasks)
		return nil
	},
}
