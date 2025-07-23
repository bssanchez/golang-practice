# Exercise 1: Basic Calculator

## Description
Implement a basic calculator that can perform basic arithmetic operations.

## Requisites
1. Implement the next functions in the `calculator.go` file:
   - `Add(a, b float64) float64` - Sum two numbers
   - `Subtract(a, b float64) float64` - Subtract two numbers
   - `Multiply(a, b float64) float64` - Multiply two numbers
   - `Divide(a, b float64) (float64, error)` - Divide two numbers, return error if b is 0
   - `Power(base, exponent float64) float64` - Calculate base raised to exponent

2. Ensure proper handling of special cases:
   - Division by zero must return an error
   - Operations with negative numbers must work correctly

## Tests
Run `go test` to verify your implementation.