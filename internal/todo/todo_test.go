package todo

import (
	"testing"
	"os"
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

// TestSaveAndLoad checks if the list can be saved to disk and fully restored
func TestSaveAndLoad(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "todo-test")
	if err != nil {
		t.Fatalf("Couldn't create temp file: %v", err)
	}

	// NOTE: defer runs this right before the function finishes to clean up the disk
	defer os.Remove(tmpfile.Name())

	list := NewList()
	list.Add("Task 1")
	list.Add("Task 2")
	list.Complete(0)

	if err := list.Save(tmpfile.Name()); err != nil {
		t.Fatalf("Failed to save list: %v", err)
	}

	loadedList := NewList()
	if err := loadedList.Load(tmpfile.Name()); err != nil {
		t.Fatalf("Failed to load list: %v", err)
	} 

	// WHY: verify that the slice length is exactly the same after loading
	if len(loadedList.Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(loadedList.Tasks))
	}
}
