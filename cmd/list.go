package cmd

import (
	"github.com/TheRedScreen64/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [<status>]",
	Short: "List all tasks",
	Long:  "List all tasks. You can filter tasks by status.",
	Example: `task-tracker list
task-tracker list todo
task-tracker list in-progress
task-tracker list done`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return task.ListTasks(args[0])
		}

		return task.ListTasks("")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
