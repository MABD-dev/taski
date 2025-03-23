package cmd

import (
	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/presentation"
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

		db := db.GetDb()
		tasks := db.List()

		if len(statuses) != 0 {
			statuses, err := stringArrayToTaskStatus(statuses)
			if err != nil {
				panic(err)
			}
			tasks = filterByStatus(tasks, statuses)
		}

		presentation.RenderTable(tasks)
		return nil
	},
}
