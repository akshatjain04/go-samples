package geometry

import (
	"fmt"
	"math"
	"os"
	"testing"
	"github.com/your_project/geometry"
)

func TestSlope(t *testing.T) {

	tests := []struct {
		name      string
		line      geometry.Line
		wantSlope float64
		wantPanic bool
	}{
		{
			name: "Horizontal Line",
			line: geometry.Line{
				P1: geometry.Point{X: 1, Y: 2},
				P2: geometry.Point{X: 3, Y: 2},
			},
			wantSlope: 0,
			wantPanic: false,
		},
		{
			name: "Vertical Line",
			line: geometry.Line{
				P1: geometry.Point{X: 2, Y: 1},
				P2: geometry.Point{X: 2, Y: 3},
			},
			wantSlope: math.Inf(1),
			wantPanic: false,
		},
		{
			name: "Diagonal Line with Negative Slope",
			line: geometry.Line{
				P1: geometry.Point{X: 1, Y: 3},
				P2: geometry.Point{X: 3, Y: 1},
			},
			wantSlope: -1,
			wantPanic: false,
		},
		{
			name: "Line with Positive Slope",
			line: geometry.Line{
				P1: geometry.Point{X: 1, Y: 1},
				P2: geometry.Point{X: 3, Y: 3},
			},
			wantSlope: 1,
			wantPanic: false,
		},
		{
			name: "Line with Same Start and End Points",
			line: geometry.Line{
				P1: geometry.Point{X: 1, Y: 1},
				P2: geometry.Point{X: 1, Y: 1},
			},
			wantPanic: true,
		},
		{
			name: "Line with Very Small Delta X",
			line: geometry.Line{
				P1: geometry.Point{X: 1, Y: 1},
				P2: geometry.Point{X: 1.0000001, Y: 2},
			},
			wantSlope: 10000000,
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Errorf("test '%s' panicked unexpectedly: %v", tt.name, r)
					} else {
						t.Logf("test '%s' expected panic: %v", tt.name, r)
					}
				}
			}()

			if tt.wantPanic {

				_ = geometry.Slope(&tt.line)
			} else {
				gotSlope := geometry.Slope(&tt.line)
				if gotSlope != tt.wantSlope {
					if math.IsInf(tt.wantSlope, 1) && math.IsInf(gotSlope, 1) {
						t.Logf("test '%s' passed, got expected positive infinity slope", tt.name)
					} else if math.IsInf(tt.wantSlope, -1) && math.IsInf(gotSlope, -1) {
						t.Logf("test '%s' passed, got expected negative infinity slope", tt.name)
					} else {
						t.Errorf("test '%s' failed, got slope = %v, want slope = %v", tt.name, gotSlope, tt.wantSlope)
					}
				} else {
					t.Logf("test '%s' passed, got expected slope = %v", tt.name, gotSlope)
				}
			}
		})
	}
}
