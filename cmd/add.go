package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add <description>",
	Short:   "Add a task",
	Example: `task-tracker add "Buy groceries"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("add called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
