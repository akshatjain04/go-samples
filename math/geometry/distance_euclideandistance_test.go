package geometry

import (
	"errors"
	"math"
	"testing"
	"github.com/your/package/geometry"
)

func TestEuclideanDistance(t *testing.T) {
	tests := []struct {
		name     string
		p1       geometry.EuclideanPoint
		p2       geometry.EuclideanPoint
		expected float64
		err      error
	}{
		{
			name:     "2D points",
			p1:       geometry.EuclideanPoint{0, 0},
			p2:       geometry.EuclideanPoint{3, 4},
			expected: 5.0,
			err:      nil,
		},
		{
			name:     "Mismatched dimensions",
			p1:       geometry.EuclideanPoint{0, 0},
			p2:       geometry.EuclideanPoint{1, 2, 3},
			expected: -1,
			err:      geometry.ErrDimMismatch,
		},
		{
			name:     "3D points",
			p1:       geometry.EuclideanPoint{0, 0, 0},
			p2:       geometry.EuclideanPoint{2, 3, 6},
			expected: 7.0,
			err:      nil,
		},
		{
			name:     "Identical points",
			p1:       geometry.EuclideanPoint{1, 2, 3},
			p2:       geometry.EuclideanPoint{1, 2, 3},
			expected: 0.0,
			err:      nil,
		},
		{
			name:     "Negative coordinates",
			p1:       geometry.EuclideanPoint{-1, -2},
			p2:       geometry.EuclideanPoint{-4, -6},
			expected: 5.0,
			err:      nil,
		},
		{
			name:     "High dimensional space",
			p1:       geometry.EuclideanPoint{1, 2, 3, 4, 5, 6, 7, 8, 9},
			p2:       geometry.EuclideanPoint{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: math.Sqrt(120),
			err:      nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			distance, err := geometry.EuclideanDistance(tc.p1, tc.p2)

			if !errors.Is(err, tc.err) {
				t.Errorf("Test %s: expected error %v, got %v", tc.name, tc.err, err)
			}

			if err == nil && distance != tc.expected {
				t.Errorf("Test %s: expected distance %v, got %v", tc.name, tc.expected, distance)
			}

			if err == nil {
				t.Logf("Test %s: passed with distance %v", tc.name, distance)
			} else {
				t.Logf("Test %s: passed with error %v", tc.name, err)
			}
		})
	}
}
