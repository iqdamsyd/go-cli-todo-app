package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var sortby string
var search string
var status string
var details bool
var complete bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos in the todo list",
	Long:  `List all todos in the todo list with their details.`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := ReadTodos()
		if err != nil {
			fmt.Println("Error reading todos:", err)
			return
		}

		if len(todos) == 0 {
			fmt.Println("No todos found.")
			return
		}

		var orderedTodos []Todo
		switch sortby {
		case "updated":
			orderedTodos = OrderTodosByLastUpdated(todos)
		case "priority":
			orderedTodos = OrderTodosByPriority(todos)
		case "deadline":
			orderedTodos = OrderTodosByDeadline(todos)
		default:
			fmt.Println("Invalid sort option. Using default sorting by last updated.")
			orderedTodos = OrderTodosByLastUpdated(todos)
		}

		if status != "" {
			orderedTodos = FilterTodosByStatus(orderedTodos, status)
		}

		completed := []Todo{}

		for _, todo := range orderedTodos {
			if todo.DeletedAt != nil {
				continue
			}

			if search != "" && !strings.Contains(strings.ToLower(todo.Description), strings.ToLower(search)) {
				continue
			}

			if todo.Completed {
				completed = append(completed, todo)
				continue
			}

			printTodo(todo, details)
		}

		if complete && len(completed) == 0 {
			fmt.Println("No completed todos found.")
			return
		}
		if complete && len(completed) > 0 {
			fmt.Println("------- COMPLETED -------")
			for _, todo := range completed {
				printTodo(todo, details)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&sortby, "sortby", "s", "updated", "Sort todos by (updated, priority, deadline)")
	listCmd.Flags().StringVarP(&search, "search", "q", "", "Search todos by description")
	listCmd.Flags().StringVarP(&status, "status", "t", "", "Filter todos by status (pending, completed)")
	listCmd.Flags().BoolVarP(&details, "details", "d", false, "Display todos with details")
	listCmd.Flags().BoolVarP(&complete, "complete", "c", false, "Display completed todos")
}

func printTodo(todo Todo, details bool) {
	status := "pending"
	if todo.Completed {
		status = "completed"
	}

	if todo.Deadline == NoDeadline {
		todo.Deadline = "-"
	}

	if details {
		printTodoDetails(todo, status)
	} else {
		printTodoSummary(todo, status)
	}
}

func printTodoDetails(todo Todo, status string) {
	fmt.Printf("ID          : %d\nDescription : %s\nPriority    : %s\nStatus      : %s\nDeadline    : %s\nCreated At  : %s\nUpdated At  : %s\n\n",
		todo.ID, todo.Description, todo.Priority, status, todo.Deadline, todo.CreatedAt.Format("2006-01-02 15:04:05"), todo.UpdatedAt.Format("2006-01-02 15:04:05"))
}

func printTodoSummary(todo Todo, status string) {
	fmt.Printf("%d | %s | %s | %s | %s\n",
		todo.ID, todo.Description, todo.Priority, status, todo.Deadline)
}
