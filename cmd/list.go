package cmd

import (
	"github.com/mabd-dev/taski/internal/domain/repos"
	"github.com/mabd-dev/taski/internal/ui"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list [-s searchTerm]...",
	Short: "List all tasks with optional search terms",
	Long:  "List all your tasks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		searchTerm, err := cmd.Flags().GetStringArray("search")
		if err != nil {
			panic(err)
		}

		tasks := repos.TasksRepo.GetAll()

		rawData := ui.TasksToKanbanRawData(tasks)
		ui.HighlightTerms(&rawData, searchTerm)
		ui.RenderRawData(rawData)

		return nil
	},
}
