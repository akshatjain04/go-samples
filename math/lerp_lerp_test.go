package math

import (
	"testing"
)

// TestLerp is a table-driven test for the Lerp function in the math package.
func TestLerp(t *testing.T) {
	tests := []struct {
		name           string
		v0, v1, factor float64
		want           float64
	}{
		{
			name:   "Basic interpolation",
			v0:     0.0,
			v1:     10.0,
			factor: 0.5,
			want:   5.0,
		},
		{
			name:   "Interpolation with t equal to 0",
			v0:     0.0,
			v1:     10.0,
			factor: 0.0,
			want:   0.0,
		},
		{
			name:   "Interpolation with t equal to 1",
			v0:     0.0,
			v1:     10.0,
			factor: 1.0,
			want:   10.0,
		},
		{
			name:   "Interpolation with t less than 0",
			v0:     0.0,
			v1:     10.0,
			factor: -0.5,
			want:   -5.0,
		},
		{
			name:   "Interpolation with t greater than 1",
			v0:     0.0,
			v1:     10.0,
			factor: 1.5,
			want:   15.0,
		},
		{
			name:   "Interpolation with negative values",
			v0:     -10.0,
			v1:     -20.0,
			factor: 0.5,
			want:   -15.0,
		},
		{
			name:   "Interpolation with v0 equal to v1",
			v0:     5.0,
			v1:     5.0,
			factor: 0.7,
			want:   5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Lerp(tt.v0, tt.v1, tt.factor)
			if got != tt.want {
				t.Errorf("Lerp(%v, %v, %v) = %v, want %v", tt.v0, tt.v1, tt.factor, got, tt.want)
			} else {
				t.Logf("Success: Lerp(%v, %v, %v) = %v", tt.v0, tt.v1, tt.factor, got)
			}
		})
	}
}
