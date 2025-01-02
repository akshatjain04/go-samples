package sort

import (
	"testing"
	"reflect"
	"sort"
)

/*
ROOST_METHOD_HASH=isSorted_fa5110ab18
ROOST_METHOD_SIG_HASH=isSorted_ef8018aa66
*/
func TestisSorted(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "Empty Array",
			arr:  []int{},
			want: true,
		},
		{
			name: "Single Element Array",
			arr:  []int{1},
			want: true,
		},
		{
			name: "Sorted Array",
			arr:  []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "Unsorted Array",
			arr:  []int{5, 4, 3, 2, 1},
			want: false,
		},
		{
			name: "Array with Duplicate Values",
			arr:  []int{1, 1, 2, 2, 3, 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSorted(tt.arr); got != tt.want {
				t.Errorf("isSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=shuffle_5769c36402
ROOST_METHOD_SIG_HASH=shuffle_9b594e5a10
*/
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
