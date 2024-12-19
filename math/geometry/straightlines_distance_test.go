package geometry

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {

	testCases := []struct {
		name     string
		pointA   Point
		pointB   Point
		expected float64
	}{
		{
			name:     "Negative coordinates",
			pointA:   Point{-3, -4},
			pointB:   Point{-1, -1},
			expected: math.Sqrt(2*2 + 3*3),
		},
		{
			name:     "Identical points",
			pointA:   Point{5, 5},
			pointB:   Point{5, 5},
			expected: 0,
		},
		{
			name:     "Large coordinate values",
			pointA:   Point{1e9, 1e9},
			pointB:   Point{-1e9, -1e9},
			expected: math.Sqrt(4e18 + 4e18),
		},
		{
			name:     "Floating-point precision",
			pointA:   Point{0.1, 0.1},
			pointB:   Point{0.2, 0.2},
			expected: math.Sqrt(0.1*0.1 + 0.1*0.1),
		},
		{
			name:     "Point at the origin and another point",
			pointA:   Point{0, 0},
			pointB:   Point{3, 4},
			expected: 5,
		},
		{
			name:     "Points on the same X-axis",
			pointA:   Point{3, 2},
			pointB:   Point{3, -2},
			expected: 4,
		},
		{
			name:     "Points on the same Y-axis",
			pointA:   Point{4, 3},
			pointB:   Point{-4, 3},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Distance(&tc.pointA, &tc.pointB)
			if math.Abs(result-tc.expected) > 1e-9 {
				t.Errorf("Distance between %v and %v was incorrect, got: %v, want: %v.", tc.pointA, tc.pointB, result, tc.expected)
			} else {
				t.Logf("Success: Distance between %v and %v was correct: %v", tc.pointA, tc.pointB, result)
			}
		})
	}
}
