package arrayops

import (
	"errors"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed numbers", []int{-5, 0, 5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.input)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"empty slice", []float64{}, 0},
		{"single element", []float64{5}, 5},
		{"integers", []float64{1, 2, 3, 4, 5}, 3},
		{"decimals", []float64{1.5, 2.5, 3.5}, 2.5},
		{"negative numbers", []float64{-1, -2, -3}, -2},
		{"mixed numbers", []float64{-5, 0, 5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Average(tt.input)
			if result != tt.expected {
				t.Errorf("Average(%v) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expected    int
		expectedErr error
	}{
		{"empty slice", []int{}, 0, errors.New("empty slice")},
		{"single element", []int{5}, 5, nil},
		{"positive numbers", []int{1, 3, 2, 5, 4}, 5, nil},
		{"negative numbers", []int{-5, -2, -10}, -2, nil},
		{"mixed numbers", []int{-5, 0, 5}, 5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Max(tt.input)
			
			// Check error
			if (err != nil) != (tt.expectedErr != nil) {
				t.Errorf("Max(%v) error = %v, wantErr %v", tt.input, err, tt.expectedErr != nil)
				return
			}
			
			// If empty slice, don't check the result
			if len(tt.input) == 0 {
				return
			}
			
			// Check result
			if result != tt.expected {
				t.Errorf("Max(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		filter   func(int) bool
		expected []int
	}{
		{
			"empty slice",
			[]int{},
			func(n int) bool { return n > 0 },
			[]int{},
		},
		{
			"filter positive",
			[]int{-2, -1, 0, 1, 2},
			func(n int) bool { return n > 0 },
			[]int{1, 2},
		},
		{
			"filter even",
			[]int{1, 2, 3, 4, 5, 6},
			func(n int) bool { return n%2 == 0 },
			[]int{2, 4, 6},
		},
		{
			"filter all",
			[]int{1, 2, 3},
			func(n int) bool { return true },
			[]int{1, 2, 3},
		},
		{
			"filter none",
			[]int{1, 2, 3},
			func(n int) bool { return false },
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.input, tt.filter)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Filter(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		mapFunc  func(int) int
		expected []int
	}{
		{
			"empty slice",
			[]int{},
			func(n int) int { return n * 2 },
			[]int{},
		},
		{
			"double",
			[]int{1, 2, 3},
			func(n int) int { return n * 2 },
			[]int{2, 4, 6},
		},
		{
			"square",
			[]int{1, 2, 3, 4},
			func(n int) int { return n * n },
			[]int{1, 4, 9, 16},
		},
		{
			"add one",
			[]int{0, 1, 2},
			func(n int) int { return n + 1 },
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.input, tt.mapFunc)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Map(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"first empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"second empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"same length", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"first longer", []int{1, 3, 5, 7}, []int{2, 4}, []int{1, 2, 3, 4, 5, 7}},
		{"second longer", []int{1, 3}, []int{2, 4, 6, 8}, []int{1, 2, 3, 4, 6, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.a, tt.b)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Merge(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}