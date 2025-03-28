package cmd

import (
	"fmt"
	"strconv"

	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
	"github.com/spf13/cobra"
)

var ViewTaskCmd = &cobra.Command{
	Use:   "view <taskNumber>",
	Short: "View task details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskNumber, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("%v is not a valid number", args[0])
		}

		task := repos.TasksRepo.Get(taskNumber)
		if task == nil {
			return fmt.Errorf("Could not find take with number=%v", taskNumber)
		}

		ui.RenderTask(*task)

		return nil
	},
}
