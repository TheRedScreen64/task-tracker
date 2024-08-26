package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Short:   "Delete a task",
	Long:    "Delete a task by specifying its id.",
	Example: "task-tracker delete 1",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("delete called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
