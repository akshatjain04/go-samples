package fibonacci

import (
	"fibonacci"
	"fmt"
	"math"
	"testing"
)

func TestMatrix(t *testing.T) {

	tests := []struct {
		name      string
		n         uint
		want      uint
		expectErr bool
	}{
		{"Test with n = 0", 0, 0, false},
		{"Test with n = 1", 1, 1, false},
		{"Test with known Fibonacci number (n = 10)", 10, 55, false},
		{"Test with a large value of n (stress test)", 1000, 0, false},
		{"Test with maximum uint value (boundary condition)", math.MaxUint32, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := fibonacci.Matrix(tt.n)

			if got != tt.want && !tt.expectErr {
				t.Errorf("Matrix(%d) = %d, want %d", tt.n, got, tt.want)
			} else if tt.expectErr && got == tt.want {
				t.Errorf("Matrix(%d) expected an error, but got correct value %d", tt.n, got)
			} else {
				t.Logf("Matrix(%d) = %d, as expected", tt.n, got)
			}
		})
	}
}

