package main

import (
	"fmt"
	"os"
	"strings"

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
	
	// WHY: os.Args[0] is always the program path, we slice [1:] to get only the words user typed
	args := os.Args[1:] // Skip the program name

	if len(args) == 0 {
		// NOTE: if user just runs the app without arguments, show current tasks and stop
		fmt.Println(todoList)
		return
	}

	// WHY: the first word after the program name is always our action (add, complete, etc.)
	command := args[0]

	if command == "add" {
		// NOTE: verify that the user actually provided some text after 'add'
		if len(args) < 2 {
			fmt.Println("Error: missing todo text")
			os.Exit(1)
		}

		// WHY: merge all separate words into a single sentence using spaces
		text := strings.Join(args[1:], " ")
		todoList.Add(text)
		saveTodos(todoList)
		fmt.Println("Added:", text)
	} else {
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: add")
		os.Exit(1)
	}
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

