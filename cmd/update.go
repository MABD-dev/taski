package cmd

import (
	"strconv"

	"github.com/mabd-dev/tasks/internal/db"
	"github.com/mabd-dev/tasks/internal/models"
	"github.com/mabd-dev/tasks/internal/renderer"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task name or description",
	Long:  "Update a task name or description by providing it's number",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// flags, name, description, status
		taskNumber, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		var name *string = nil
		if cmd.Flags().Changed("name") {
			nameStr, err := cmd.Flags().GetString("name")
			if err != nil {
				panic(err)
			}
			name = &nameStr
		}

		var description *string = nil
		if cmd.Flags().Changed("description") {
			descStr, err := cmd.Flags().GetString("description")
			if err != nil {
				panic(err)
			}
			description = &descStr
		}

		var status *models.TaskStatus = nil
		if cmd.Flags().Changed("status") {
			statusStr, err := cmd.Flags().GetString("status")
			if err != nil {
				panic(err)
			}
			s, err := models.TaskStatusStrToStatus(statusStr)
			if err != nil {
				panic(err)
			}
			status = &s
		}

		db := db.GetDb()
		err = db.Update(taskNumber, name, description, status)
		if err != nil {
			panic("could not find task")
		}

		renderer.RenderTable(db.List())
	},
}
