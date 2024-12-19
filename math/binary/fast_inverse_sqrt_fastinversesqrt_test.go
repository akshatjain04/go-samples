package binary

import (
	"binary"
	"math"
	"testing"
)

func TestFastInverseSqrt(t *testing.T) {

	testCases := []struct {
		description    string
		input          float32
		want           float32
		errorMargin    float32
		expectInfinity bool
		expectNaN      bool
	}{
		{
			description:    "Validate correct inverse square root for a positive number",
			input:          4.0,
			want:           1 / 2.0,
			errorMargin:    0.0001,
			expectInfinity: false,
			expectNaN:      false,
		},
		{
			description:    "Validate result for zero input",
			input:          0.0,
			want:           float32(math.Inf(1)),
			errorMargin:    0,
			expectInfinity: true,
			expectNaN:      false,
		},
		{
			description:    "Validate result for a negative number",
			input:          -4.0,
			want:           float32(math.NaN()),
			errorMargin:    0,
			expectInfinity: false,
			expectNaN:      true,
		},
		{
			description:    "Validate precision of the result",
			input:          10.0,
			want:           1 / math.Sqrt(10.0),
			errorMargin:    0.0001,
			expectInfinity: false,
			expectNaN:      false,
		},
		{
			description:    "Validate behavior for very large input values",
			input:          float32(math.MaxFloat32),
			want:           1 / float32(math.Sqrt(math.MaxFloat32)),
			errorMargin:    0.0001,
			expectInfinity: false,
			expectNaN:      false,
		},
		{
			description:    "Validate behavior for very small input values",
			input:          float32(math.SmallestNonzeroFloat32),
			want:           1 / math.Sqrt(float64(math.SmallestNonzeroFloat32)),
			errorMargin:    0.0001,
			expectInfinity: false,
			expectNaN:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := binary.FastInverseSqrt(tc.input)

			switch {
			case tc.expectInfinity:
				if !math.IsInf(float64(got), 1) {
					t.Errorf("FastInverseSqrt(%v): expected positive infinity, got %v", tc.input, got)
				}
			case tc.expectNaN:
				if !math.IsNaN(float64(got)) {
					t.Errorf("FastInverseSqrt(%v): expected NaN, got %v", tc.input, got)
				}
			default:
				if math.Abs(float64(got-tc.want)) > float64(tc.errorMargin) {
					t.Errorf("FastInverseSqrt(%v): expected %v, got %v", tc.input, tc.want, got)
				}
			}

			t.Logf("Passed: %s", tc.description)
		})
	}
}

