package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
	Long:  "Delete a task given it's number",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskNumber := args[0]
		fmt.Printf("Deleting a task with number=%v", taskNumber)
	},
}
