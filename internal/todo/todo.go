package todo

import (
	"errors"
	"fmt"
	"encoding/json"
	"os"
)

// Task represents a single todo item
type Task struct {
	Text string 
	Done bool
}

// NewTask creates and returns a new Task instance
func NewTask(text string) Task {
	// WHY: no '*' here because task is tiny, low memory
	return Task{
		Text: text,
		Done: false,
	}
} 

// List represents a collection of tasks
type List struct {
	// NOTE: don't need to initialize with make(), append() works with nil slices anyway
	Tasks []Task
}

// NewList initializes and returns a pointer to a new empty List
func NewList() *List {
	// WHY: returns '&' (pointer) so we always change the EXACT SAME list everywhere
	return &List{}
}

// Add appends a new task to the list
func (l *List) Add(text string) {
	// WHY: need '*' in (l *List) otherwise Go creates a copy and original list stays empty
	task := NewTask(text)
	l.Tasks = append(l.Tasks, task)
}

// Complete marks a task as done by its index
func (l *List) Complete(index int) error {
	// NOTE: check if index makes sense, otherwise the app will just crash
	if index < 0 || index >= len(l.Tasks) {
		return errors.New("task index out of range") // Go rule: lowercase, no dot at the end
	}

	l.Tasks[index].Done = true

	return nil
}

// String formats the entire task list into a user-friendly string
func (l *List) String() string {
	if len(l.Tasks) == 0 {
		return "No tasks in the todo list"
	}

	result := "Todo List:\n"
	// WHY: used '*' here just to be consistent with other methods (Add, Complete).
	for i, task := range l.Tasks {
		status := " "

		if task.Done {
			status = "✅"
		}

		result += fmt.Sprintf("%d. [%s] %s\n", i+1, status, task.Text)
	}

	return result
}

// Save writes the todo list to a file in JSON format
func (l *List) Save(filename string) error {
	// WHY: convert struct to JSON bytes
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}

	// NOTE: 0644 means Owner: Read+Write (6), Group: Read (4), Others: Read (4)
	return os.WriteFile(filename, data, 0644)
}

// Load reads a todo list from a file
func (l *List) Load(filename string) error {
	// WHY: read raw bytes from the hard drive
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// WHY: 'l' is already a pointer here, so we pass it directly without '&'
	// Unmarshal needs the exact memory address to unpack JSON data into it
	return json.Unmarshal(data, l)
}
