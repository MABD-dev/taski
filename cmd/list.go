package cmd

import (
	"github.com/mabd-dev/tasks/internal/db"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all your tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		db := db.GetDb()
		db.List()
	},
}
