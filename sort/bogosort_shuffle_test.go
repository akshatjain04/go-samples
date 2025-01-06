package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext // For running tests and subtests.
}
func TestShuffle(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name          string
		input         []int
		expectedLen   int
		expectedError bool
	}{
		{
			name:        "Valid input with positive integers",
			input:       []int{1, 2, 3, 4, 5},
			expectedLen: 5,
		},
		{
			name:        "Valid input with negative integers",
			input:       []int{-1, -2, -3, -4, -5},
			expectedLen: 5,
		},
		{
			name:        "Valid input with floating-point numbers",
			input:       []int{1.1, 2.2, 3.3, 4.4, 5.5},
			expectedLen: 5,
		},
		{
			name:        "Empty slice input",
			input:       []int{},
			expectedLen: 0,
		},
		{
			name:        "Single element slice input",
			input:       []int{42},
			expectedLen: 1,
		},
		{
			name:        "Large input slice",
			input:       make([]int, 10000),
			expectedLen: 10000,
		},
		{
			name:        "Input slice with duplicates",
			input:       []int{1, 1, 2, 2, 3, 3},
			expectedLen: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			sort.Shuffle(tt.input)

			if len(tt.input) != tt.expectedLen {
				t.Errorf("Shuffled slice has incorrect length: got %d, want %d", len(tt.input), tt.expectedLen)
			}

			if !reflect.DeepEqual(original, tt.input) {
				t.Logf("Shuffled slice has the same elements as the original")
			}

			if isDifferentOrder(original, tt.input) {
				t.Logf("Shuffled slice is in a different order than the original")
			} else if tt.expectedLen > 1 {
				t.Errorf("Shuffled slice is in the same order as the original")
			}
		})
	}
}
func isDifferentOrder(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}
