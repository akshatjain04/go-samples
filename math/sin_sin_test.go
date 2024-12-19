package math

import (
	"math"
	"testing"
)

// TestSin is a test function for the Sin function in the math package.
func TestSin(t *testing.T) {
	// Table-driven tests for the Sin function
	tests := []struct {
		name         string
		input        float64
		want         float64
		wantTolerance float64
	}{
		{"Zero value", 0, 0, 0},
		{"Positive value", math.Pi / 2, 1, 1e-14},
		{"Negative value", -math.Pi / 2, -1, 1e-14},
		{"Overflow value", math.MaxFloat64, 0, 0}, // The expected result is hard to predict due to the nature of floating-point arithmetic
		{"Periodic value", 2 * math.Pi, 0, 0},
		{"Precision check", math.Pi / 4, math.Sqrt(2) / 2, 1e-14},
		// Scenario 7 is skipped as Go is statically typed and will not compile with incorrect types
		{"Extreme values (near limits)", math.SmallestNonzeroFloat64, 0, 0}, // The expected result is hard to predict due to the nature of floating-point arithmetic
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sin(tt.input)
			if math.Abs(got-tt.want) > tt.wantTolerance {
				t.Errorf("Sin(%v) = %v, want %v", tt.input, got, tt.want)
			} else {
				t.Logf("Sin(%v) = %v, passed", tt.input, got)
			}
		})
	}
}
