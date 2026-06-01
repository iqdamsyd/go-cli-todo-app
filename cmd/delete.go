package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [todo ID]",
	Short: "Delete a todo from the todo list",
	Long:  `Delete a todo from the todo list by providing its ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the todo ID to delete.")
			return
		}

		todoID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid todo ID:", err)
			return
		}

		err = DeleteTodo(todoID)
		if err != nil {
			fmt.Println("Error deleting todo:", err)
			return
		}

		fmt.Printf("Todo with ID %d deleted successfully.\n", todoID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
