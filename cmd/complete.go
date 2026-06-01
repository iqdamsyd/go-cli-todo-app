package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [todo ID]",
	Short: "Mark a todo as completed",
	Long:  `Mark a todo as completed by providing its ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the todo ID to mark as completed.")
			return
		}

		todoID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid todo ID:", err)
			return
		}

		err = MarkTodoAsCompleted(todoID)
		if err != nil {
			fmt.Println("Error marking todo as completed:", err)
			return
		}

		fmt.Printf("Todo with ID %d marked as completed.\n", todoID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
