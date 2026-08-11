package todo

import (
	"testing"
)

// TestNewTask checks if a single task is created with correct defaults
func TestNewTask(t *testing.T) {
	task := NewTask("Learn Go")

	if task.Text != "Learn Go" {
		t.Errorf("Expected text to be 'Learn Go', got '%s'", task.Text)
	}

	if task.Done {
		// NOTE: fresh tasks must always be false
		t.Error("New task shouldn't be marked as done")
	}
}

// TestAddTask checks if tasks are correctly appended to the list
func TestAddTask(t *testing.T) {
	list := NewList()
	list.Add("Feed the cat")

	if len(list.Tasks) != 1 {
		t.Errorf("Expected 1 item in list, got %d", len(list.Tasks))
	}

	if list.Tasks[0].Text != "Feed the cat" {
		t.Errorf("Expected text to be 'Feed the cat', got '%s'", list.Tasks[0].Text)
	}
}

// TestCompleteTask checks normal completion and out-of-bounds errors
func TestCompleteTask(t *testing.T) {
	list := NewList()
	list.Add("Write code on Go")

	// 1. Test normal case (index 0 exists)
	err := list.Complete(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !list.Tasks[0].Done {
		t.Error("Expected task to be marked as done")
	}

	// 2. Test error case (index 1 does not exist)
	err = list.Complete(1)
	if err == nil {
		// WHY: 'err == nil' means no error happened, but we EXPECTED an error here
		t.Error("Expected error when completing non-existent task")
	}
}
