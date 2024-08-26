package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [<status>]",
	Short: "List all tasks",
	Long:  "List all tasks. You can filter tasks by status.",
	Example: `task-tracker list todo
    task-tracker list in-progress
    task-tracker list done`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("list called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
