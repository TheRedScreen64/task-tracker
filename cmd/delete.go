package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/TheRedScreen64/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Short:   "Delete a task",
	Long:    "Delete a task by specifying its id.",
	Example: "task-tracker delete 1",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("task id is required")
		}

		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("%s is not a valid id", args[0])
		}

		return task.DeleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
