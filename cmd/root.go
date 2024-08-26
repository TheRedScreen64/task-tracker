package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "A simple cli based task tracker",
	Long: `A simple cli tool for managing tasks.
	
Features:
 - Add, Update, and Delete tasks
 - Mark a task as in progress or done
 - List all tasks
 - List all tasks that are done
 - List all tasks that are not done
 - List all tasks that are in progress
 
Source Code: https://github.com/TheRedScreen64/task-tracker`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
