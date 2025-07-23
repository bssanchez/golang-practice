# Ejercicio 3: Operaciones con Arrays

## Descripción
Implementa funciones para realizar operaciones comunes con arrays y slices en Go.

## Requisitos
1. Implementa las siguientes funciones en el archivo `arrayops.go`:
   - `Sum(numbers []int) int` - Suma todos los elementos de un slice
   - `Average(numbers []float64) float64` - Calcula el promedio de un slice
   - `Max(numbers []int) (int, error)` - Encuentra el valor máximo (error si slice vacío)
   - `Filter(numbers []int, f func(int) bool) []int` - Filtra elementos según una función
   - `Map(numbers []int, f func(int) int) []int` - Aplica una función a cada elemento
   - `Merge(a, b []int) []int` - Combina dos slices alternando sus elementos

2. Consideraciones:
   - Las funciones deben manejar slices vacíos correctamente
   - Para `Merge`, si los slices tienen diferente longitud, añade los elementos restantes al final

## Pruebas
Ejecuta `go test` para verificar tu implementación.