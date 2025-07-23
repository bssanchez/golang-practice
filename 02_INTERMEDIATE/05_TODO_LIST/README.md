# Exercise 5: Todo List

## Description
Implement a task management system (TODO list) using structs and methods in Go.

## Requirements
1. Create a `Task` struct in the `todolist.go` file with the following fields:
   - `ID` (int)
   - `Title` (string)
   - `Description` (string)
   - `Done` (bool)
   - `CreatedAt` (time.Time)
   - `CompletedAt` (time.Time, can be null)

2. Create a `TodoList` struct with the following fields:
   - `Tasks` (slice of Task)
   - `nextID` (int, to generate unique IDs)

3. Implement the following methods for `TodoList`:
   - `NewTodoList() *TodoList` - Constructor
   - `AddTask(title, description string) int` - Adds a task and returns its ID
   - `DeleteTask(id int) error` - Deletes a task by ID
   - `CompleteTask(id int) error` - Marks a task as completed
   - `GetTask(id int) (Task, error)` - Gets a task by ID
   - `ListTasks(onlyPending bool) []Task` - Lists all tasks or only pending ones
   - `SearchTasks(query string) []Task` - Searches tasks by title or description

4. Considerations:
   - Tasks must have unique and auto-incremental IDs
   - When completing a task, the current date must be recorded in `CompletedAt`
   - Searches must be case-insensitive

## Tests
Run `go test` to verify your implementation.