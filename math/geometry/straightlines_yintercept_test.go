package geometry

import (
	"fmt"
	"math"
	"os"
	"testing"
	"github.com/your_project/geometry"
)

func TestYIntercept(t *testing.T) {

	testCases := []struct {
		name        string
		point       geometry.Point
		slope       float64
		expected    float64
		expectError bool
	}{
		{
			name:     "Point on the Y-axis",
			point:    geometry.Point{X: 0, Y: 7},
			slope:    2,
			expected: 7,
		},
		{
			name:     "Positive slope",
			point:    geometry.Point{X: 2, Y: 5},
			slope:    3,
			expected: -1,
		},
		{
			name:     "Negative slope",
			point:    geometry.Point{X: 1, Y: 4},
			slope:    -2,
			expected: 6,
		},
		{
			name:     "Zero slope (horizontal line)",
			point:    geometry.Point{X: 5, Y: 3},
			slope:    0,
			expected: 3,
		},
		{
			name:        "Vertical line (undefined slope)",
			point:       geometry.Point{X: 3, Y: 2},
			slope:       math.Inf(1),
			expected:    math.Inf(1),
			expectError: false,
		},
		{
			name:     "Point with negative coordinates",
			point:    geometry.Point{X: -4, Y: -2},
			slope:    1.5,
			expected: 4,
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			yIntercept := geometry.YIntercept(&tc.point, tc.slope)

			if tc.expectError {
				if !math.IsInf(yIntercept, 1) {
					t.Errorf("Expected infinite Y-intercept for vertical line, got %v", yIntercept)
				}
			} else if yIntercept != tc.expected {
				t.Errorf("Expected Y-intercept to be %v, got %v", tc.expected, yIntercept)
			} else {
				t.Logf("YIntercept() = %v, want %v", yIntercept, tc.expected)
			}
		})
	}

	w.Close()
	os.Stdout = oldStdout

	var buf []byte
	fmt.Fscanf(r, "%s", &buf)
	r.Close()

	t.Logf("Captured output: %s", string(buf))
}
