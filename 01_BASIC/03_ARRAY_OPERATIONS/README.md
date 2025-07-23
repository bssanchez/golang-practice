# Exercise 3: Array Operations

## Description
Implement functions to perform common operations with arrays and slices in Go.

## Requirements
1. Implement the following functions in the `arrayops.go` file:
   - `Sum(numbers []int) int` - Sum all elements in a slice
   - `Average(numbers []float64) float64` - Calculate the average of a slice
   - `Max(numbers []int) (int, error)` - Find the maximum value (error if slice is empty)
   - `Filter(numbers []int, f func(int) bool) []int` - Filter elements according to a function
   - `Map(numbers []int, f func(int) int) []int` - Apply a function to each element
   - `Merge(a, b []int) []int` - Combine two slices by alternating their elements

2. Considerations:
   - Functions must handle empty slices correctly
   - For `Merge`, if slices have different lengths, add the remaining elements at the end

## Tests
Run `go test` to verify your implementation.