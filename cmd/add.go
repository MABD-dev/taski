package cmd

import (
	"github.com/mabd-dev/tasks/internal/db"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [task name] [task description]",
	Short: "Add new task",
	Long:  "Add new task to the list with default completion value to false",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		description := ""
		if len(args) > 1 {
			description = args[1]
		}

		db := db.GetDb()
		db.Add(name, description)
		db.List()
	},
}
