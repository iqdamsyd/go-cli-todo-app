package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var deadline string
var priority string

var addCmd = &cobra.Command{
	Use:   "add [todo description]",
	Short: "Add a new todo to the todo list",
	Long:  `Add a new todo to the todo list with the specified description.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := Todo{
			ID:          0,
			Description: fmt.Sprintf("%s", strings.Join(args, " ")),
			Priority:    priority,
			Deadline:    deadline,
			Completed:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}

		todos, err := ReadTodos()
		if err != nil {
			fmt.Println("Error reading todos:", err)
			return
		}

		todo.ID = len(todos) + 1
		todos = append(todos, todo)

		err = WriteTodo(todos)
		if err != nil {
			fmt.Println("Error writing todos:", err)
			return
		}

		fmt.Printf("Todo added: %s\n", todo.Description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&priority, "priority", "p", "medium", "Set the priority of the todo (low, medium, high)")
	addCmd.Flags().StringVarP(&deadline, "deadline", "d", "", "Set the deadline for the todo (YYYY-MM-DD)")
}
