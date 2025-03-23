package cmd

import (
	"strconv"

	"github.com/mabd-dev/taski/internal/data/db"
	"github.com/mabd-dev/taski/internal/presentation"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
	Long: `delete a task given it's number or list of numbers like:
$ ./tasks delete 1 2 3 4 5
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		taskNumbers := []int{}

		for _, t := range args {
			taskNumber, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			taskNumbers = append(taskNumbers, taskNumber)
		}

		db := db.GetDb()
		err := db.DeleteAll(taskNumbers)
		if err != nil {
			return err
		}

		presentation.RenderTable(db.List())
		return nil
	},
}
