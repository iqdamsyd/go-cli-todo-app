package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var todoFile string

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exePath)

	todoFile = filepath.Join(exeDir, "todos.json")
}

func ReadTodos() ([]Todo, error) {
	// Implement the logic to read todos from the JSON file and return them as a slice of Todo structs.
	// Handle any errors that may occur during file reading or JSON unmarshalling.

	var todos []Todo
	fileData, err := os.ReadFile(todoFile)
	if err != nil {
		return nil, err
	}

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &todos)
		if err != nil {
			return nil, err
		}
	}

	return todos, nil
}

func WriteTodo(todos []Todo) error {
	// Implement the logic to write the slice of Todo structs to the JSON file.
	// Handle any errors that may occur during JSON marshalling or file writing.

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(todoFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTodo(id int, updatedTodo Todo) error {
	// Implement the logic to update a todo with the given ID in the JSON file.
	// Read the existing todos, find the todo with the specified ID, update its fields, and write the updated todos back to the file.
	// Handle any errors that may occur during this process.
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	found := false

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	err = WriteTodo(todos)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(id int) error {
	// Implement the logic to delete a todo with the given ID from the JSON file.
	// Read the existing todos, filter out the todo with the specified ID, and write the remaining todos back to the file.
	// Handle any errors that may occur during this process.
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	var updatedTodos []Todo
	found := false

	for _, todo := range todos {
		if todo.ID == id {
			found = true
			continue
		}
		updatedTodos = append(updatedTodos, todo)
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	err = WriteTodo(updatedTodos)
	if err != nil {
		return err
	}

	return nil
}

func MarkTodoAsCompleted(id int) error {
	// Implement the logic to mark a todo as completed by updating its Completed field to true.
	// Use the UpdateTodo function to update the todo in the JSON file.
	// Handle any errors that may occur during this process.

	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	var updatedTodo Todo
	found := false

	for _, todo := range todos {
		if todo.ID == id {
			todo.Completed = true
			todo.UpdatedAt = time.Now()
			updatedTodo = todo
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	err = UpdateTodo(id, updatedTodo)
	if err != nil {
		return err
	}

	return nil
}

func OrderTodosByLastUpdated(todos []Todo) []Todo {
	// Implement the logic to order todos by their last updated time.
	// You can define a custom sorting function to sort the todos based on their UpdatedAt field.
	// Return the ordered slice of Todo structs.

	sortedTodos := make([]Todo, len(todos))
	copy(sortedTodos, todos)

	for i := 0; i < len(sortedTodos)-1; i++ {
		for j := 0; j < len(sortedTodos)-i-1; j++ {
			if sortedTodos[j].UpdatedAt.Before(sortedTodos[j+1].UpdatedAt) {
				sortedTodos[j], sortedTodos[j+1] = sortedTodos[j+1], sortedTodos[j]
			}
		}
	}

	return sortedTodos
}

func OrderTodosByPriority(todos []Todo) []Todo {
	// Implement the logic to order todos by their priority (e.g., high, medium, low).
	// You can define a custom sorting function to sort the todos based on their Priority field.
	// Return the ordered slice of Todo structs.

	priorityOrder := map[string]int{
		"high":   1,
		"medium": 2,
		"low":    3,
	}

	sortedTodos := make([]Todo, len(todos))
	copy(sortedTodos, todos)

	for i := 0; i < len(sortedTodos)-1; i++ {
		for j := 0; j < len(sortedTodos)-i-1; j++ {
			if priorityOrder[sortedTodos[j].Priority] > priorityOrder[sortedTodos[j+1].Priority] {
				sortedTodos[j], sortedTodos[j+1] = sortedTodos[j+1], sortedTodos[j]
			} else {
				// If priorities are the same, sort by deadline
				if priorityOrder[sortedTodos[j].Priority] == priorityOrder[sortedTodos[j+1].Priority] {
					if sortedTodos[j].Deadline > sortedTodos[j+1].Deadline {
						sortedTodos[j], sortedTodos[j+1] = sortedTodos[j+1], sortedTodos[j]
					}
				}
			}
		}
	}

	return sortedTodos
}

func OrderTodosByDeadline(todos []Todo) []Todo {
	// Implement the logic to order todos by their deadline.
	// You can define a custom sorting function to sort the todos based on their Deadline field.
	// Return the ordered slice of Todo structs.

	sortedTodos := make([]Todo, len(todos))
	copy(sortedTodos, todos)

	for i := 0; i < len(sortedTodos)-1; i++ {
		for j := 0; j < len(sortedTodos)-i-1; j++ {
			if sortedTodos[j].Deadline > sortedTodos[j+1].Deadline {
				sortedTodos[j], sortedTodos[j+1] = sortedTodos[j+1], sortedTodos[j]
			}
		}
	}

	return sortedTodos
}

func FilterTodosByStatus(todos []Todo, status string) []Todo {
	// Implement the logic to filter todos based on their completion status (e.g., completed, pending).
	// Return a slice of Todo structs that match the specified status.

	var filteredTodos []Todo

	for _, todo := range todos {
		if status == "completed" && todo.Completed {
			filteredTodos = append(filteredTodos, todo)
		} else if status == "pending" && !todo.Completed {
			filteredTodos = append(filteredTodos, todo)
		}
	}

	return filteredTodos
}
