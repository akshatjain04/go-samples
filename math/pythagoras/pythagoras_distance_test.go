package pythagoras

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {

	testCases := []struct {
		name             string
		vectorA          Vector
		vectorB          Vector
		expectedDistance float64
	}{
		{
			name:             "Equal Points",
			vectorA:          Vector{0, 0, 0},
			vectorB:          Vector{0, 0, 0},
			expectedDistance: 0.0,
		},
		{
			name:             "Positive Axis Distance",
			vectorA:          Vector{0, 0, 0},
			vectorB:          Vector{5, 0, 0},
			expectedDistance: 5.0,
		},
		{
			name:             "Negative Axis Distance",
			vectorA:          Vector{0, 0, 0},
			vectorB:          Vector{-5, 0, 0},
			expectedDistance: 5.0,
		},
		{
			name:             "Diagonal Distance",
			vectorA:          Vector{0, 0, 0},
			vectorB:          Vector{3, 4, 5},
			expectedDistance: math.Sqrt(50),
		},
		{
			name:             "Large Values",
			vectorA:          Vector{1e9, 1e9, 1e9},
			vectorB:          Vector{1e9 + 3, 1e9 - 4, 1e9 + 5},
			expectedDistance: math.Sqrt(50),
		},
		{
			name:             "Small Values",
			vectorA:          Vector{0.0001, 0.0001, 0.0001},
			vectorB:          Vector{0.0004, 0.0005, 0.0006},
			expectedDistance: math.Sqrt(0.0003*0.0003 + 0.0004*0.0004 + 0.0005*0.0005),
		},
		{
			name:             "Zero Distance",
			vectorA:          Vector{0, 0, 0},
			vectorB:          Vector{0, 0, 0},
			expectedDistance: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actualDistance := Distance(tc.vectorA, tc.vectorB)

			if actualDistance != tc.expectedDistance {
				t.Errorf("Distance between %v and %v was incorrect, got: %v, want: %v",
					tc.vectorA, tc.vectorB, actualDistance, tc.expectedDistance)
			} else {
				t.Logf("Successfully calculated distance for test case: %s", tc.name)
			}
		})
	}
}
