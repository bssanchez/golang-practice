# Ejercicio 1: Calculadora Básica

## Descripción
Implementa una calculadora básica que pueda realizar operaciones matemáticas simples.

## Requisitos
1. Implementa las siguientes funciones en el archivo `calculator.go`:
   - `Add(a, b float64) float64` - Suma dos números
   - `Subtract(a, b float64) float64` - Resta dos números
   - `Multiply(a, b float64) float64` - Multiplica dos números
   - `Divide(a, b float64) (float64, error)` - Divide dos números, devuelve error si b es 0
   - `Power(base, exponent float64) float64` - Calcula base elevado a exponent

2. Asegúrate de manejar correctamente los casos especiales:
   - División por cero debe devolver un error
   - Operaciones con números negativos deben funcionar correctamente

## Pruebas
Ejecuta `go test` para verificar tu implementación.