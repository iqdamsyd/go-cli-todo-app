package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [todo ID] [new description]",
	Short: "Update the description of a todo",
	Long:  `Update the description of a todo by providing its ID and the new description.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please provide the todo ID and the new description.")
			return
		}

		todoID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid todo ID:", err)
			return
		}

		todos, err := ReadTodos()
		if err != nil {
			fmt.Println("Error reading todos:", err)
			return
		}

		var todo Todo
		for _, t := range todos {
			if t.ID == todoID {
				todo = t
				break
			}
		}

		if todo.ID == 0 {
			fmt.Printf("Todo with ID %d not found.\n", todoID)
			return
		}

		todo.Description = fmt.Sprintf("%s", strings.Join(args[1:], " "))
		todo.UpdatedAt = time.Now()

		if priority != "" {
			todo.Priority = priority
		}

		if deadline != "" {
			todo.Deadline = deadline
		}

		err = UpdateTodo(todoID, todo)
		if err != nil {
			fmt.Println("Error updating todo description:", err)
			return
		}

		fmt.Printf("Todo with ID %d updated successfully.\n", todoID)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&priority, "priority", "p", "", "Update the priority of the todo (low, medium, high)")
	updateCmd.Flags().StringVarP(&deadline, "deadline", "d", "", "Update the deadline for the todo (YYYY-MM-DD)")
}
