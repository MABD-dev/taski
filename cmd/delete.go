package cmd

import (
	"fmt"
	"strconv"

	"github.com/mabd-dev/tasks/internal/db"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
	Long:  "Delete a task given it's number",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		taskNumber, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("Deleting a task with number=%v\n", taskNumber)

		db := db.GetDb()
		err = db.Delete(taskNumber)
		if err != nil {
			return err
		}

		db.List()
		return nil
	},
}
