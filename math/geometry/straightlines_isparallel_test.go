package geometry

import (
	"math"
	"testing"
	"github.com/TheAlgorithms/Go/math/geometry"
)

func TestIsParallel(t *testing.T) {

	const tolerance = 1e-9

	isApproxEqual := func(a, b float64) bool {
		return math.Abs(a-b) <= tolerance
	}

	tests := []struct {
		name     string
		line1    geometry.Line
		line2    geometry.Line
		expected bool
	}{
		{
			name:     "Lines with different slopes are not parallel",
			line1:    geometry.Line{P1: geometry.Point{0, 0}, P2: geometry.Point{1, 1}},
			line2:    geometry.Line{P1: geometry.Point{0, 0}, P2: geometry.Point{1, 2}},
			expected: false,
		},
		{
			name:     "Lines with the same slope are parallel",
			line1:    geometry.Line{P1: geometry.Point{0, 0}, P2: geometry.Point{1, 1}},
			line2:    geometry.Line{P1: geometry.Point{1, 1}, P2: geometry.Point{2, 2}},
			expected: true,
		},
		{
			name:     "Vertical lines are parallel",
			line1:    geometry.Line{P1: geometry.Point{1, 0}, P2: geometry.Point{1, 1}},
			line2:    geometry.Line{P1: geometry.Point{2, 0}, P2: geometry.Point{2, 1}},
			expected: true,
		},
		{
			name:     "Horizontal lines are parallel",
			line1:    geometry.Line{P1: geometry.Point{0, 1}, P2: geometry.Point{1, 1}},
			line2:    geometry.Line{P1: geometry.Point{0, 2}, P2: geometry.Point{1, 2}},
			expected: true,
		},
		{
			name:     "Line compared with itself is parallel",
			line1:    geometry.Line{P1: geometry.Point{0, 0}, P2: geometry.Point{1, 1}},
			line2:    geometry.Line{P1: geometry.Point{0, 0}, P2: geometry.Point{1, 1}},
			expected: true,
		},
		{
			name:     "Handling of floating-point precision",
			line1:    geometry.Line{P1: geometry.Point{0, 0}, P2: geometry.Point{1, 1}},
			line2:    geometry.Line{P1: geometry.Point{0, 0}, P2: geometry.Point{1, 1 + tolerance}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := geometry.IsParallel(&tt.line1, &tt.line2)
			if !isApproxEqual(geometry.Slope(&tt.line1), geometry.Slope(&tt.line2)) {
				actual = false
			}
			if actual != tt.expected {
				t.Errorf("IsParallel(%v, %v) = %v, expected %v", tt.line1, tt.line2, actual, tt.expected)
			} else {
				t.Logf("IsParallel(%v, %v) = %v, as expected", tt.line1, tt.line2, actual)
			}
		})
	}
}
