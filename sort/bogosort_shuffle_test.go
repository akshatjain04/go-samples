package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Testshuffle(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name     string
		input    []float64
		validate func([]float64, []float64) bool
	}{
		{
			name:  "Valid input with positive integers",
			input: []float64{1, 2, 3, 4, 5},
			validate: func(original, shuffled []float64) bool {
				return reflect.DeepEqual(original, shuffled) && !reflect.DeepEqual(original, shuffled)
			},
		},
		{
			name:  "Valid input with negative integers",
			input: []float64{-1, -2, -3, -4, -5},
			validate: func(original, shuffled []float64) bool {
				return reflect.DeepEqual(original, shuffled) && !reflect.DeepEqual(original, shuffled)
			},
		},
		{
			name:  "Valid input with floating-point numbers",
			input: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			validate: func(original, shuffled []float64) bool {
				return reflect.DeepEqual(original, shuffled) && !reflect.DeepEqual(original, shuffled)
			},
		},
		{
			name:  "Empty slice input",
			input: []float64{},
			validate: func(original, shuffled []float64) bool {
				return len(shuffled) == 0
			},
		},
		{
			name:  "Single element slice input",
			input: []float64{42},
			validate: func(original, shuffled []float64) bool {
				return reflect.DeepEqual(original, shuffled)
			},
		},
		{
			name:  "Large input slice",
			input: make([]float64, 10000),
			validate: func(original, shuffled []float64) bool {
				return reflect.DeepEqual(original, shuffled) && !reflect.DeepEqual(original, shuffled)
			},
		},
		{
			name:  "Input slice with duplicates",
			input: []float64{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			validate: func(original, shuffled []float64) bool {
				return reflect.DeepEqual(original, shuffled) && !reflect.DeepEqual(original, shuffled)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			original := make([]float64, len(tc.input))
			copy(original, tc.input)

			shuffle(tc.input)

			if !tc.validate(original, tc.input) {
				t.Errorf("Failed %s: shuffled slice does not match validation criteria", tc.name)
			} else {
				t.Logf("Passed %s", tc.name)
			}
		})
	}
}


