package cmd

import (
	"strconv"

	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <task number>...",
	Short: "Delete a task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		taskNumbers := []int{}

		for _, t := range args {
			taskNumber, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			taskNumbers = append(taskNumbers, taskNumber)
		}

		err := repos.TasksRepo.Delete(taskNumbers...)
		if err != nil {
			return err
		}

		ui.RenderKanbanBoard(repos.TasksRepo.GetAll())
		return nil
	},
}
