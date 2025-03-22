/*
Copyright © 2025 MABD-dev <mabd.universe@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage your daily tasks from your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		ListCmd.Run(cmd, args)
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

	rootCmd.AddCommand(DeleteCmd)
}
