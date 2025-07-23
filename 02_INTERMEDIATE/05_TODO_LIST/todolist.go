package todolist

import (
	"errors"
	"strings"
	"time"
)

// Task representa una tarea en la lista de tareas
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// TodoList representa una lista de tareas
type TodoList struct {
	Tasks  []Task
	nextID int
}

// NewTodoList crea una nueva lista de tareas
func NewTodoList() *TodoList {
	// TODO: Implementar esta función
	return nil
}

// AddTask añade una nueva tarea a la lista
// Devuelve el ID de la tarea creada
func (l *TodoList) AddTask(title, description string) int {
	// TODO: Implementar esta función
	return 0
}

// DeleteTask elimina una tarea por su ID
// Devuelve error si la tarea no existe
func (l *TodoList) DeleteTask(id int) error {
	// TODO: Implementar esta función
	return nil
}

// CompleteTask marca una tarea como completada
// Devuelve error si la tarea no existe o ya está completada
func (l *TodoList) CompleteTask(id int) error {
	// TODO: Implementar esta función
	return nil
}

// GetTask obtiene una tarea por su ID
// Devuelve error si la tarea no existe
func (l *TodoList) GetTask(id int) (Task, error) {
	// TODO: Implementar esta función
	return Task{}, nil
}

// ListTasks devuelve todas las tareas o solo las pendientes
// Si onlyPending es true, solo devuelve las tareas no completadas
func (l *TodoList) ListTasks(onlyPending bool) []Task {
	// TODO: Implementar esta función
	return nil
}

// SearchTasks busca tareas que contengan el query en el título o descripción
// La búsqueda es case-insensitive
func (l *TodoList) SearchTasks(query string) []Task {
	// TODO: Implementar esta función
	return nil
}