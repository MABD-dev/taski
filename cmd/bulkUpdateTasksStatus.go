package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mabd-dev/taski/internal/domain/models"
	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
	"github.com/spf13/cobra"
)

var BulkUpdateTasksStatus = &cobra.Command{
	Use:   "status <status> <taskNumber>...",
	Short: "Bulk update status to multiple tasks ",
	Long: `Given new status, set that value to all provided task numbers

Operation starts after checking that all task numbers are valid.
`,
	Args: cobra.MinimumNArgs(1), //arg is new project name
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return errors.New("At least provide one task number to update")
		}

		status, err := models.TaskStatusStrToStatus(args[0])
		if err != nil {
			return err
		}

		taskNumbersMap := map[int]int{}
		tasks := []*models.Task{}

		// checking if all are valid numbers
		for _, s := range args[1:] {
			n, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("%v is not a valid number", s)
			}
			if _, exists := taskNumbersMap[n]; exists {
				continue
			}

			// Get all tasks and make sure they exist
			task := repos.TasksRepo.Get(n)
			if task == nil {
				return fmt.Errorf("task not found with number=%v", n)
			}
			tasks = append(tasks, task)
		}

		// update all tasks
		for _, task := range tasks {
			(*task).Status = status
			err := repos.TasksRepo.Update(task.Number, *task)
			if err != nil {
				return err
			}
		}

		ui.ClearTerminal()
		ui.RenderKanbanBoard(repos.TasksRepo.GetAll())

		return nil
	},
}
