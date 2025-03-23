/*
Copyright © 2025 MABD-dev <mabd.universe@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/mabd-dev/taski/repl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taski",
	Short: "Start REPL session",
	RunE: func(cmd *cobra.Command, args []string) error {
		repl.StartRepl()
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(ListCmd)
	ListCmd.PersistentFlags().StringArrayP("status", "s", []string{}, "Filter tasks by status. Not case sensitive")

	rootCmd.AddCommand(AddCmd)
	AddCmd.PersistentFlags().StringP("status", "s", "todo", "Add status to task. options[\"todo\", \"inprogress\", \"done\"]")
	AddCmd.PersistentFlags().StringP("description", "d", "", "Add description to task")

	rootCmd.AddCommand(UpdateCmd)
	UpdateCmd.PersistentFlags().StringP("name", "n", "", "Update name of the task")
	UpdateCmd.PersistentFlags().StringP("description", "d", "", "Add description to task")
	UpdateCmd.PersistentFlags().StringP("status", "s", "", "Update status of the task. Old values will be removed!")

	rootCmd.AddCommand(DeleteCmd)
}
