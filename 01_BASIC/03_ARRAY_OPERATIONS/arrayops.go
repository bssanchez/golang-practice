package arrayops

import (
	"errors"
)

// Sum calculates the sum of all elements in a slice
func Sum(numbers []int) int {
	// TODO: Implement this function
	return 0
}

// Average calculates the average of elements in a slice
func Average(numbers []float64) float64 {
	// TODO: Implement this function
	return 0
}

// Max finds the maximum value in a slice
// Returns an error if the slice is empty
func Max(numbers []int) (int, error) {
	// TODO: Implement this function
	return 0, nil
}

// Filter filters the elements of a slice according to a predicate function
func Filter(numbers []int, f func(int) bool) []int {
	// TODO: Implement this function
	return nil
}

// Map applies a function to each element of a slice
func Map(numbers []int, f func(int) int) []int {
	// TODO: Implement this function
	return nil
}

// Merge combines two slices by alternating their elements
// If the slices have different lengths, it adds the remaining elements at the end
func Merge(a, b []int) []int {
	// TODO: Implement this function
	return nil
}