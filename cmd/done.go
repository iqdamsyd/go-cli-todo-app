package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [todo ID]",
	Short: "Mark a todo as done",
	Long:  `Mark a todo as done by providing its ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the todo ID to mark as done.")
			return
		}

		todoID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid todo ID:", err)
			return
		}

		err = MarkTodoAsDone(todoID)
		if err != nil {
			fmt.Println("Error marking todo as done:", err)
			return
		}

		fmt.Printf("Todo with ID %d marked as done.\n", todoID)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
