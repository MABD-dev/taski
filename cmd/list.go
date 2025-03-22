package cmd

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mabd-dev/tasks/internal/db"
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

		table := table.New(os.Stdout)
		table.SetHeaders("#", "Name", "Description", "Status", "Created At")

		for _, task := range tasks {
			table.AddRow(strconv.Itoa(task.Number), task.Name, task.Description, task.Status.ToString(), task.CreatedAt.Format(time.RFC1123))
		}
		table.Render()
	},
}
