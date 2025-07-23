package todolist

import (
	"strings"
	"testing"
	"time"
)

func TestNewTodoList(t *testing.T) {
	list := NewTodoList()
	
	if list == nil {
		t.Fatal("NewTodoList() returned nil")
	}
	
	if list.Tasks == nil {
		t.Error("Tasks slice should be initialized")
	}
	
	if len(list.Tasks) != 0 {
		t.Errorf("New list should have 0 tasks, got %d", len(list.Tasks))
	}
}

func TestAddTask(t *testing.T) {
	list := NewTodoList()
	
	id1 := list.AddTask("Buy groceries", "Milk, eggs, bread")
	if id1 != 1 {
		t.Errorf("First task should have ID 1, got %d", id1)
	}
	
	id2 := list.AddTask("Clean house", "Vacuum and dust")
	if id2 != 2 {
		t.Errorf("Second task should have ID 2, got %d", id2)
	}
	
	if len(list.Tasks) != 2 {
		t.Errorf("List should have 2 tasks, got %d", len(list.Tasks))
	}
	
	// Check task properties
	task, err := list.GetTask(1)
	if err != nil {
		t.Fatalf("Unexpected error getting task: %v", err)
	}
	
	if task.Title != "Buy groceries" {
		t.Errorf("Expected title 'Buy groceries', got '%s'", task.Title)
	}
	
	if task.Description != "Milk, eggs, bread" {
		t.Errorf("Expected description 'Milk, eggs, bread', got '%s'", task.Description)
	}
	
	if task.Done {
		t.Error("New task should not be marked as done")
	}
	
	// Check creation time
	now := time.Now()
	if task.CreatedAt.After(now) || task.CreatedAt.Before(now.Add(-time.Minute)) {
		t.Errorf("Task creation time should be recent, got %v", task.CreatedAt)
	}
	
	// CompletedAt should be zero time
	if !task.CompletedAt.IsZero() {
		t.Error("CompletedAt should be zero for new task")
	}
}

func TestDeleteTask(t *testing.T) {
	list := NewTodoList()
	
	// Add tasks
	list.AddTask("Task 1", "Description 1")
	list.AddTask("Task 2", "Description 2")
	list.AddTask("Task 3", "Description 3")
	
	// Delete middle task
	err := list.DeleteTask(2)
	if err != nil {
		t.Errorf("Unexpected error deleting task: %v", err)
	}
	
	if len(list.Tasks) != 2 {
		t.Errorf("List should have 2 tasks after deletion, got %d", len(list.Tasks))
	}
	
	// Try to get deleted task
	_, err = list.GetTask(2)
	if err == nil {
		t.Error("Expected error getting deleted task, got nil")
	}
	
	// Delete non-existent task
	err = list.DeleteTask(99)
	if err == nil {
		t.Error("Expected error deleting non-existent task, got nil")
	}
	
	// Check remaining tasks
	task1, _ := list.GetTask(1)
	if task1.Title != "Task 1" {
		t.Errorf("Expected task 1 title to be 'Task 1', got '%s'", task1.Title)
	}
	
	task3, _ := list.GetTask(3)
	if task3.Title != "Task 3" {
		t.Errorf("Expected task 3 title to be 'Task 3', got '%s'", task3.Title)
	}
}

func TestCompleteTask(t *testing.T) {
	list := NewTodoList()
	
	// Add a task
	id := list.AddTask("Test task", "Test description")
	
	// Complete the task
	err := list.CompleteTask(id)
	if err != nil {
		t.Errorf("Unexpected error completing task: %v", err)
	}
	
	// Check task status
	task, _ := list.GetTask(id)
	if !task.Done {
		t.Error("Task should be marked as done")
	}
	
	// Check completion time
	now := time.Now()
	if task.CompletedAt.After(now) || task.CompletedAt.Before(now.Add(-time.Minute)) {
		t.Errorf("Task completion time should be recent, got %v", task.CompletedAt)
	}
	
	// Try to complete already completed task
	err = list.CompleteTask(id)
	if err == nil {
		t.Error("Expected error completing already completed task, got nil")
	}
	
	// Try to complete non-existent task
	err = list.CompleteTask(99)
	if err == nil {
		t.Error("Expected error completing non-existent task, got nil")
	}
}

func TestListTasks(t *testing.T) {
	list := NewTodoList()
	
	// Add tasks
	list.AddTask("Task 1", "Description 1")
	list.AddTask("Task 2", "Description 2")
	list.AddTask("Task 3", "Description 3")
	
	// Complete one task
	list.CompleteTask(2)
	
	// List all tasks
	allTasks := list.ListTasks(false)
	if len(allTasks) != 3 {
		t.Errorf("Expected 3 tasks in total, got %d", len(allTasks))
	}
	
	// List only pending tasks
	pendingTasks := list.ListTasks(true)
	if len(pendingTasks) != 2 {
		t.Errorf("Expected 2 pending tasks, got %d", len(pendingTasks))
	}
	
	// Check that completed tasks are not in pending list
	for _, task := range pendingTasks {
		if task.Done {
			t.Errorf("Found completed task in pending list: %+v", task)
		}
	}
}

func TestSearchTasks(t *testing.T) {
	list := NewTodoList()
	
	// Add tasks
	list.AddTask("Buy groceries", "Milk, eggs, bread")
	list.AddTask("Clean house", "Vacuum and dust")
	list.AddTask("Buy new phone", "Check prices online")
	
	// Search by title
	results := list.SearchTasks("buy")
	if len(results) != 2 {
		t.Errorf("Expected 2 results for 'buy', got %d", len(results))
	}
	
	// Search by description
	results = list.SearchTasks("milk")
	if len(results) != 1 {
		t.Errorf("Expected 1 result for 'milk', got %d", len(results))
	}
	
	// Case insensitive search
	results = list.SearchTasks("HOUSE")
	if len(results) != 1 {
		t.Errorf("Expected 1 result for 'HOUSE', got %d", len(results))
	}
	
	// Search with no results
	results = list.SearchTasks("nonexistent")
	if len(results) != 0 {
		t.Errorf("Expected 0 results for 'nonexistent', got %d", len(results))
	}
}