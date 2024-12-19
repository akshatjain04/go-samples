package geometry

import (
	"math"
	"testing"
)

func TestIsPerpendicular(t *testing.T) {

	testCases := []struct {
		name     string
		line1    Line
		line2    Line
		expected bool
	}{
		{
			name:     "Non-inverse slopes",
			line1:    Line{Point{0, 0}, Point{1, 1}},
			line2:    Line{Point{0, 0}, Point{1, 2}},
			expected: false,
		},
		{
			name:     "Perpendicular positive and negative slopes",
			line1:    Line{Point{0, 0}, Point{1, 1}},
			line2:    Line{Point{0, 0}, Point{1, -1}},
			expected: true,
		},
		{
			name:     "Horizontal and vertical lines",
			line1:    Line{Point{0, 0}, Point{1, 0}},
			line2:    Line{Point{0, 0}, Point{0, 1}},
			expected: true,
		},
		{
			name:     "Identical lines",
			line1:    Line{Point{0, 0}, Point{1, 1}},
			line2:    Line{Point{0, 0}, Point{1, 1}},
			expected: false,
		},
		{
			name:     "Zero slope and finite non-zero slope",
			line1:    Line{Point{0, 0}, Point{1, 0}},
			line2:    Line{Point{0, 0}, Point{1, 2}},
			expected: false,
		},
		{
			name:     "Both lines vertical",
			line1:    Line{Point{0, 0}, Point{0, 1}},
			line2:    Line{Point{1, 0}, Point{1, 1}},
			expected: false,
		},
		{
			name:     "Very close to perpendicular slopes",
			line1:    Line{Point{0, 0}, Point{1, 1}},
			line2:    Line{Point{0, 0}, Point{1, -1.0000001}},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPerpendicular(&tc.line1, &tc.line2)
			if result != tc.expected {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("%s: succeeded, lines are %v as expected", tc.name, result)
			}
		})
	}
}
