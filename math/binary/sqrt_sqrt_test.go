package binary

import (
	"github.com/TheAlgorithms/Go/math/binary"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {

	var tests = []struct {
		name      string
		input     float32
		want      float32
		expectErr bool
	}{
		{
			name:      "NegativeNumber",
			input:     -4.0,
			want:      0,
			expectErr: true,
		},
		{
			name:      "VeryLargeNumber",
			input:     1e+30,
			want:      1e+15,
			expectErr: false,
		},
		{
			name:      "SmallestPositiveFloat32",
			input:     math.SmallestNonzeroFloat32,
			want:      math.Sqrt(math.SmallestNonzeroFloat32),
			expectErr: false,
		},
		{
			name:      "PositiveNonSquareNumber",
			input:     2,
			want:      math.Sqrt(2),
			expectErr: false,
		},
		{
			name:      "Float32MaxValue",
			input:     math.MaxFloat32,
			want:      math.Sqrt(math.MaxFloat32),
			expectErr: false,
		},
		{
			name:      "PrecisionLimits",
			input:     1.234567,
			want:      math.Sqrt(1.234567),
			expectErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := binary.Sqrt(tc.input)
			if tc.expectErr {
				if !math.IsNaN(float64(got)) {
					t.Errorf("Sqrt(%v) = %v; want NaN", tc.input, got)
				} else {
					t.Logf("Sqrt(%v) correctly returned NaN", tc.input)
				}
			} else {
				if math.Abs(float64(got)-float64(tc.want)) > 1e-5 {
					t.Errorf("Sqrt(%v) = %v; want %v", tc.input, got, tc.want)
				} else {
					t.Logf("Sqrt(%v) = %v; want %v", tc.input, got, tc.want)
				}
			}
		})
	}
}
