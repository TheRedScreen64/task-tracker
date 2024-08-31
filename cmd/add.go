package cmd

import (
	"errors"

	"github.com/TheRedScreen64/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add <description>",
	Short:   "Add a task",
	Example: `task-tracker add "Buy groceries"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("task description is required")
		}

		return task.AddTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
