package sort

import (
	"math/rand"
	"testing"
	"reflect"
	"sort"
)


func TestShuffle(t *testing.T) {

	testCases := []struct {
		name        string
		input       []int
		expectPanic bool
	}{
		{
			name:        "Valid Shuffle Test",
			input:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectPanic: false,
		},
		{
			name:        "Empty Slice Test",
			input:       []int{},
			expectPanic: false,
		},
		{
			name:        "Single Element Slice Test",
			input:       []int{1},
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			original := make([]int, len(tc.input))
			copy(original, tc.input)

			shuffle(tc.input)

			if tc.expectPanic {
				if !reflect.DeepEqual(tc.input, original) {
					t.Errorf("shuffle() = %v, want = %v", tc.input, original)
				}
			} else {

				if reflect.DeepEqual(tc.input, original) {
					t.Errorf("shuffle() = %v, want a different order of elements", tc.input)
				}

				sort.Ints(tc.input)
				sort.Ints(original)
				if !reflect.DeepEqual(tc.input, original) {
					t.Errorf("shuffle() = %v, want = %v", tc.input, original)
				}
			}
		})
	}
}

