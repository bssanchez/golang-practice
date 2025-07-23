# Ejercicio 5: Lista de Tareas

## Descripción
Implementa un sistema de gestión de tareas (TODO list) utilizando structs y métodos en Go.

## Requisitos
1. Crea un struct `Task` en el archivo `todolist.go` con los siguientes campos:
   - `ID` (int)
   - `Title` (string)
   - `Description` (string)
   - `Done` (bool)
   - `CreatedAt` (time.Time)
   - `CompletedAt` (time.Time, puede ser nula)

2. Crea un struct `TodoList` con los siguientes campos:
   - `Tasks` (slice de Task)
   - `nextID` (int, para generar IDs únicos)

3. Implementa los siguientes métodos para `TodoList`:
   - `NewTodoList() *TodoList` - Constructor
   - `AddTask(title, description string) int` - Añade una tarea y devuelve su ID
   - `DeleteTask(id int) error` - Elimina una tarea por ID
   - `CompleteTask(id int) error` - Marca una tarea como completada
   - `GetTask(id int) (Task, error)` - Obtiene una tarea por ID
   - `ListTasks(onlyPending bool) []Task` - Lista todas las tareas o solo las pendientes
   - `SearchTasks(query string) []Task` - Busca tareas por título o descripción

4. Consideraciones:
   - Las tareas deben tener IDs únicos y autoincrementales
   - Al completar una tarea, se debe registrar la fecha actual en `CompletedAt`
   - Las búsquedas deben ser case-insensitive

## Pruebas
Ejecuta `go test` para verificar tu implementación.