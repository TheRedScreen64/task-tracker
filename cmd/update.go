package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/TheRedScreen64/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update <task-id> <new-description>",
	Short:   "Update a task",
	Long:    "Update a task by specifying its id and the new description.",
	Example: `task-tracker update 1 "Cook dinner"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("task id and the new description is required")
		}

		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("%s is not a valid id", args[0])
		}

		return task.UpdateTaskDesc(id, args[1])
	},
}

var markInProgressCmd = &cobra.Command{
	Use:     "mark-in-progress <task-id>",
	Short:   "Mark a task as in-progress",
	Example: `task-tracker mark-in-progress 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("task id is required")
		}

		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("%s is not a valid id", args[0])
		}

		return task.UpdateTaskStatus(id, task.TASK_STATUS_IN_PROGRESS)
	},
}

var markDoneCmd = &cobra.Command{
	Use:     "mark-done <task-id>",
	Short:   "Mark a task as done",
	Example: `task-tracker mark-done 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("task id is required")
		}

		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("%s is not a valid id", args[0])
		}

		return task.UpdateTaskStatus(id, task.TASK_STATUS_DONE)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)
}
