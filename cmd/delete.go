package cmd

import (
	"strconv"

	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
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

		err := repos.TasksRepo.DeleteAll(taskNumbers)
		if err != nil {
			return err
		}

		ui.RenderKanbanBoard(repos.TasksRepo.List())
		return nil
	},
}
