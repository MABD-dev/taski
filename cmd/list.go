package cmd

import (
	"github.com/mabd-dev/tasks/internal/db"
	"github.com/mabd-dev/tasks/internal/renderer"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all your tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
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

		renderer.RenderTable(tasks)
	},
}
