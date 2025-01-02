package sort

import "testing"

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
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
