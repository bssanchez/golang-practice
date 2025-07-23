# Ejercicio 6: Procesador de Archivos

## Descripción
Implementa un procesador de archivos que pueda leer, escribir y analizar archivos de texto.

## Requisitos
1. Implementa las siguientes funciones en el archivo `fileprocessor.go`:
   - `ReadFile(filename string) ([]string, error)` - Lee un archivo y devuelve sus líneas
   - `WriteFile(filename string, lines []string) error` - Escribe líneas a un archivo
   - `CountWords(filename string) (int, error)` - Cuenta palabras en un archivo
   - `FindPattern(filename, pattern string) ([]string, error)` - Encuentra líneas que coinciden con un patrón
   - `ReplaceInFile(filename, old, new string) (int, error)` - Reemplaza texto y devuelve el número de reemplazos

2. Consideraciones:
   - Maneja correctamente los errores de archivo (no existe, permisos, etc.)
   - Para `FindPattern`, usa expresiones regulares
   - Para `CountWords`, considera que las palabras están separadas por espacios, tabs o saltos de línea

## Pruebas
Ejecuta `go test` para verificar tu implementación.

## Nota
Para las pruebas, se crearán archivos temporales en el directorio actual.