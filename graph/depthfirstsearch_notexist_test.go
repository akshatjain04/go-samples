package graph

import (
	"testing"
)

func TestNotExist(t *testing.T) {
	type args struct {
		target int
		slice  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Target element exists in the slice",
			args: args{
				target: 5,
				slice:  []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "Target element does not exist in the slice",
			args: args{
				target: 6,
				slice:  []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "Empty slice",
			args: args{
				target: 5,
				slice:  []int{},
			},
			want: true,
		},
		{
			name: "Target element is the first element in the slice",
			args: args{
				target: 1,
				slice:  []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "Target element is the last element in the slice",
			args: args{
				target: 5,
				slice:  []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "Target element is the only element in the slice",
			args: args{
				target: 1,
				slice:  []int{1},
			},
			want: false,
		},
		{
			name: "Target element is a negative integer",
			args: args{
				target: -5,
				slice:  []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "Target element is a large integer",
			args: args{
				target: 100,
				slice:  []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotExist(tt.args.target, tt.args.slice); got != tt.want {
				t.Errorf("NotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
