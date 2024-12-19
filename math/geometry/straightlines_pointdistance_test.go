package geometry

import (
	"geometry"
	"math"
	"testing"
)

func TestPointDistance(t *testing.T) {

	tests := []struct {
		name     string
		point    geometry.Point
		equation [3]float64
		want     float64
		wantErr  bool
	}{
		{
			name:     "Point on the line",
			point:    geometry.Point{X: 2, Y: -3},
			equation: [3]float64{1, 1, -5},
			want:     0,
		},
		{
			name:     "Negative coefficients in the equation",
			point:    geometry.Point{X: 3, Y: 4},
			equation: [3]float64{-1, -1, 5},
			want:     4.242640687119285,
		},
		{
			name:     "Point with negative coordinates",
			point:    geometry.Point{X: -3, Y: -4},
			equation: [3]float64{2, 3, 6},
			want:     6,
		},
		{
			name:     "Vertical line equation",
			point:    geometry.Point{X: 3, Y: 0},
			equation: [3]float64{1, 0, -3},
			want:     0,
		},
		{
			name:     "Horizontal line equation",
			point:    geometry.Point{X: 0, Y: 4},
			equation: [3]float64{0, 1, -4},
			want:     0,
		},
		{
			name:     "Zero coefficients in the equation",
			point:    geometry.Point{X: 1, Y: 1},
			equation: [3]float64{0, 0, 0},
			wantErr:  true,
		},
		{
			name:     "Large numbers for coordinates and coefficients",
			point:    geometry.Point{X: 1e9, Y: 1e9},
			equation: [3]float64{1e9, 1e9, -1e18},
			want:     0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := geometry.PointDistance(&tc.point, tc.equation)
			if tc.wantErr {
				if got != 0 {
					t.Errorf("PointDistance() with zero coefficients should return 0; got %v", got)
				}
			} else if !almostEqual(got, tc.want) {
				t.Errorf("PointDistance() = %v, want %v", got, tc.want)
			} else {
				t.Logf("Success: %s", tc.name)
			}
		})
	}
}
func almostEqual(a, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}

