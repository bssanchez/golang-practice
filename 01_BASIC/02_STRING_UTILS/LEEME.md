# Ejercicio 2: Utilidades de Strings

## Descripción
Implementa una serie de funciones para manipular strings en Go.

## Requisitos
1. Implementa las siguientes funciones en el archivo `stringutils.go`:
   - `Reverse(s string) string` - Invierte un string
   - `IsPalindrome(s string) bool` - Verifica si un string es un palíndromo
   - `CountOccurrences(s, substr string) int` - Cuenta las ocurrencias de un substring
   - `ToTitleCase(s string) string` - Convierte la primera letra de cada palabra a mayúscula
   - `RemoveDuplicateChars(s string) string` - Elimina caracteres duplicados

2. Consideraciones:
   - Las funciones deben manejar strings vacíos correctamente
   - Para `IsPalindrome`, ignora mayúsculas/minúsculas y espacios
   - Para `ToTitleCase`, considera que las palabras están separadas por espacios

## Pruebas
Ejecuta `go test` para verificar tu implementación.