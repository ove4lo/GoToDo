package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ove4lo/GoToDo/internal/todo"
)

const (
	todoFile = "todos.json"
)

func main() {
	// WHY: flag.Bool returns a pointer (*bool) because it allocates memory before parsing
	interactiveFlag := flag.Bool("i", false, "Run in interactive mode")

	// WHY: flag.Bool defines standard '-h' help flag
	helpFlag := flag.Bool("h", false, "Show help information")

	// NOTE: flag.Parse() extracts flags, flag.Args() leaves only clean arguments
	flag.Parse()
	args := flag.Args()

	// Show help if the -h flag is provided
	if *helpFlag {
		printHelp()
		return
	}

	// Load existing todos
	todoList := todo.NewList()

	// NOTE: check if file exists first, so Load() won't break on the very first run
	if _, err := os.Stat(todoFile); err == nil {
		if err := todoList.Load(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, "Error loading todos:", err)
			os.Exit(1)
		}
	}

	// WHY: if pointer value is true, start the endless loop and exit main immediately
	if *interactiveFlag {
		runInteractive(todoList)
		return
	}

	// Handle subcommands
	if len(args) == 0 {
		// Default action when no command is provided
		fmt.Println(todoList)
		return
	}

	// WHY: the first word after the program name is always our action (add, complete, etc.)
	command := args[0]

	switch command {
	case "add":
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

	case "list":
		// WHY: just print the list, fmt will format it with checkboxes automatically
		fmt.Println(todoList)

	case "complete":
		// NOTE: check if user provided the task number to complete
		if len(args) < 2 {
			fmt.Println("Error: missing item number")
			os.Exit(1)
		}

		// WHY: convert string argument (like "1") into a real integer number
		num, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid item number:", args[1])
			os.Exit(1)
		}

		if err := todoList.Complete(num - 1); err != nil {
			fmt.Fprintln(os.Stderr, "Error completing todo:", err)
			os.Exit(1)
		}

		saveTodos(todoList)
		fmt.Println("Marked item as completed")

	case "help":
		printHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run 'todo help' or 'todo -h' for usage information")
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

// runInteractive starts an interactive shell loop for managing todos
func runInteractive(list *todo.List) {
	// WHY: bufio.Scanner reads the whole line with spaces, unlike fmt.Scan
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n" + list.String())
		fmt.Println("\nCommands:")
		fmt.Println("	add <text>	    - Add a new todo")
		fmt.Println("	complete <n>	- Mark item n as completed")
		fmt.Println("	quit			- Exit the program")
		fmt.Print("\n> ")

		// NOTE: waits for user input. Returns false if terminal closes
		if !scanner.Scan() {
			break // Exits the loop, but we don't know why yet
		}

		input := scanner.Text()
		// WHY: SplitN splits ONLY by the first space into 2 parts: [command, full_text]
		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]

		switch cmd {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Error: missing todo text")
				continue  // WHY: jumps back to the top of the for loop to ask input again
			}

			list.Add(parts[1])
			saveTodos(list)
			fmt.Println("Added:", parts[1])

		case "complete":
			if len(parts) < 2 {
				fmt.Println("Error: missing item number")
				continue
			}

			num, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error: invalid item number")
				continue
			}

			if err := list.Complete(num - 1); err != nil {
				fmt.Println("Error:", err)
				continue
			}

			saveTodos(list)
			fmt.Println("Marked item as completed")
		
		case "help":
			fmt.Println("\nAvailable commands:")
			fmt.Println("	add <text>		- Add a new todo")
			fmt.Println("	list			- List all todos")
			fmt.Println("	complete <n>	- Mark item n as completed")
			fmt.Println("	help			- Show this help message")
			fmt.Println("	quit			- Exit the program")

		case "quit", "exit":
			return // stops the entire function, exiting the loop

		default:
			fmt.Println("Unknown command:", cmd)
		}
	}

	// WHY: check if the loop stopped because of a real system error or just EOF
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1) // Stop the app with error code if something actually broke
	}
}

// printHelp outputs the full usage guide for the CLI application
func printHelp() {
	// WHY: backticks allow creating readable multi-line text without using '\n'
	helpText := `
	Todo - A simple command line todo manager
	
	Usage:
	todo [command] [arguments]
	todo[flags]
	
	Commands:
		add <text>		Add a new todo item
		list			List all todo items
		complete <n>	Mark item n as completed
		help			Show this help message
		
	Flags:
		-h				Show this help message
		-i				Run in interactive mode
	
	Examples:
		todo add "Learn Go testing"
		todo list
		todo complete 2
		todo -i
	`

	fmt.Println(helpText)
}