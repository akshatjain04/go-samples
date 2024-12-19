package geometry

import (
	"fmt"
	"geometry"
	"math"
	"testing"
)

func TestSection(t *testing.T) {

	type testCase struct {
		p1, p2 geometry.Point
		r      float64
		want   geometry.Point
		name   string
	}

	tests := []testCase{
		{
			name: "Equal distance between points",
			p1:   geometry.Point{X: 0, Y: 0},
			p2:   geometry.Point{X: 2, Y: 2},
			r:    1,
			want: geometry.Point{X: 1, Y: 1},
		},
		{
			name: "Ratio greater than 1",
			p1:   geometry.Point{X: 0, Y: 0},
			p2:   geometry.Point{X: 3, Y: 3},
			r:    2,
			want: geometry.Point{X: 2, Y: 2},
		},
		{
			name: "Ratio less than 1",
			p1:   geometry.Point{X: 0, Y: 0},
			p2:   geometry.Point{X: 3, Y: 3},
			r:    0.5,
			want: geometry.Point{X: 1, Y: 1},
		},
		{
			name: "Ratio is zero",
			p1:   geometry.Point{X: 1, Y: 1},
			p2:   geometry.Point{X: 3, Y: 3},
			r:    0,
			want: geometry.Point{X: 1, Y: 1},
		},
		{
			name: "Negative ratio",
			p1:   geometry.Point{X: 1, Y: 1},
			p2:   geometry.Point{X: 3, Y: 3},
			r:    -1,
			want: geometry.Point{X: 0, Y: 0},
		},
		{
			name: "Points with negative coordinates",
			p1:   geometry.Point{X: -1, Y: -1},
			p2:   geometry.Point{X: 1, Y: 1},
			r:    1,
			want: geometry.Point{X: 0, Y: 0},
		},
		{
			name: "Points are the same",
			p1:   geometry.Point{X: 2, Y: 2},
			p2:   geometry.Point{X: 2, Y: 2},
			r:    0.5,
			want: geometry.Point{X: 2, Y: 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			got := geometry.Section(&tc.p1, &tc.p2, tc.r)

			if got.X != tc.want.X || got.Y != tc.want.Y {
				t.Errorf("Section(%v, %v, %v) = %v, want %v", tc.p1, tc.p2, tc.r, got, tc.want)
			} else {
				t.Logf("Section(%v, %v, %v) = %v, as expected", tc.p1, tc.p2, tc.r, got)
			}
		})
	}
}

