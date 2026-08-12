package main

import (
	"fmt"
	"os"

	"github.com/ove4lo/GoToDo/internal/todo"
)

const (
	todoFile = "todos.json"
)

func main() {
	todoList := todo.NewList()

	// NOTE: check if file exists first, so Load() won't break on the very first run
	if _, err := os.Stat(todoFile); err == nil {
		if err := todoList.Load(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, "Error loading todos:", err)
			os.Exit(1)
		}
	}

	// WHY: fmt.Println automatically uses our String() method to format the output
	fmt.Println(todoList)
}

// saveTodos writes the todo list to the JSON file.
func saveTodos(list *todo.List) {
	// WHY: use '*' to work with our real list, not a temporary copy
	if err := list.Save(todoFile); err != nil {
		// NOTE: print error to the screen and completely stop the app right here
		fmt.Fprintln(os.Stderr, "Error saving todos:", err)
		os.Exit(1)
	}
}

