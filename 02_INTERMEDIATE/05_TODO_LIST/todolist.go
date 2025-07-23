package todolist

import (
	"errors"
	"strings"
	"time"
)

// Task represents a task in the todo list
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// TodoList represents a list of tasks
type TodoList struct {
	Tasks  []Task
	nextID int
}

// NewTodoList creates a new todo list
func NewTodoList() *TodoList {
	// TODO: Implement this function
	return nil
}

// AddTask adds a new task to the list
// Returns the ID of the created task
func (l *TodoList) AddTask(title, description string) int {
	// TODO: Implement this function
	return 0
}

// DeleteTask deletes a task by its ID
// Returns error if the task doesn't exist
func (l *TodoList) DeleteTask(id int) error {
	// TODO: Implement this function
	return nil
}

// CompleteTask marks a task as completed
// Returns error if the task doesn't exist or is already completed
func (l *TodoList) CompleteTask(id int) error {
	// TODO: Implement this function
	return nil
}

// GetTask gets a task by its ID
// Returns error if the task doesn't exist
func (l *TodoList) GetTask(id int) (Task, error) {
	// TODO: Implement this function
	return Task{}, nil
}

// ListTasks returns all tasks or only pending ones
// If onlyPending is true, only returns uncompleted tasks
func (l *TodoList) ListTasks(onlyPending bool) []Task {
	// TODO: Implement this function
	return nil
}

// SearchTasks searches for tasks that contain the query in title or description
// The search is case-insensitive
func (l *TodoList) SearchTasks(query string) []Task {
	// TODO: Implement this function
	return nil
}