package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update <task-id> <new-description>",
	Short:   "Update a task",
	Long:    "Update a task by specifying its id and the new description.",
	Example: `task-tracker update 1 "Cook dinner"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("update called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
