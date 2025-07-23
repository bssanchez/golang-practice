package calculator

import (
	"errors"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed numbers", -5, 8, 3},
		{"zeros", 0, 0, 0},
		{"decimals", 2.5, 3.5, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed numbers", -5, 3, -8},
		{"zeros", 0, 0, 0},
		{"decimals", 5.5, 2.5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -5, -3, 15},
		{"mixed numbers", -5, 3, -15},
		{"zeros", 0, 5, 0},
		{"decimals", 2.5, 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectedErr error
	}{
		{"positive numbers", 6, 3, 2, nil},
		{"negative numbers", -6, -3, 2, nil},
		{"mixed numbers", -6, 3, -2, nil},
		{"division by zero", 5, 0, 0, errors.New("division by zero")},
		{"decimals", 5, 2, 2.5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			
			// Check error
			if (err != nil) != (tt.expectedErr != nil) {
				t.Errorf("Divide(%f, %f) error = %v, wantErr %v", tt.a, tt.b, err, tt.expectedErr != nil)
				return
			}
			
			// If division by zero, don't check the result
			if tt.b == 0 {
				return
			}
			
			// Check result
			if result != tt.expected {
				t.Errorf("Divide(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name           string
		base, exponent float64
		expected       float64
	}{
		{"positive base and exponent", 2, 3, 8},
		{"negative base, even exponent", -2, 2, 4},
		{"negative base, odd exponent", -2, 3, -8},
		{"zero base", 0, 5, 0},
		{"zero exponent", 5, 0, 1},
		{"fractional exponent", 4, 0.5, 2}, // square root of 4
		{"negative exponent", 2, -2, 0.25}, // 1/(2^2)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.base, tt.exponent)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Power(%f, %f) = %f; want %f", tt.base, tt.exponent, result, tt.expected)
			}
		})
	}
}