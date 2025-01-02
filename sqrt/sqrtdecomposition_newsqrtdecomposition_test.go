package sqrt

import (
	"math"
	"testing"
	"reflect"
)

func TestNewSqrtDecomposition(t *testing.T) {

	scenarios := []struct {
		name               string
		elements           []int
		querySingleElement func(int) int
		unionQ             func(int, int) int
		updateQ            func(int, int, int) int
		want               *SqrtDecomposition[int, int]
	}{
		{
			name:               "Valid Scenario with Integer Elements and Sum Query",
			elements:           []int{1, 2, 3, 4, 5},
			querySingleElement: func(element int) int { return element },
			unionQ:             func(q1, q2 int) int { return q1 + q2 },
			updateQ:            func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
			want: &SqrtDecomposition[int, int]{
				querySingleElement: func(element int) int { return element },
				unionQ:             func(q1, q2 int) int { return q1 + q2 },
				updateQ:            func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
				elements:           []int{1, 2, 3, 4, 5},
				blocks:             make([]int, int(math.Ceil(math.Sqrt(5)))),
				blockSize:          uint64(math.Sqrt(5)),
			},
		},
		{
			name:               "Empty Elements Slice",
			elements:           []int{},
			querySingleElement: func(element int) int { return element },
			unionQ:             func(q1, q2 int) int { return q1 + q2 },
			updateQ:            func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
			want: &SqrtDecomposition[int, int]{
				querySingleElement: func(element int) int { return element },
				unionQ:             func(q1, q2 int) int { return q1 + q2 },
				updateQ:            func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
				elements:           []int{},
				blocks:             []int{},
				blockSize:          0,
			},
		},
		{
			name:               "Large Elements Slice",
			elements:           make([]int, 10000),
			querySingleElement: func(element int) int { return element },
			unionQ:             func(q1, q2 int) int { return q1 + q2 },
			updateQ:            func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
			want: &SqrtDecomposition[int, int]{
				querySingleElement: func(element int) int { return element },
				unionQ:             func(q1, q2 int) int { return q1 + q2 },
				updateQ:            func(oldQ, oldE, newE int) int { return oldQ - oldE + newE },
				elements:           make([]int, 10000),
				blocks:             make([]int, 100),
				blockSize:          100,
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			got := NewSqrtDecomposition(s.elements, s.querySingleElement, s.unionQ, s.updateQ)
			if !reflect.DeepEqual(got, s.want) {
				t.Errorf("NewSqrtDecomposition() = %v, want %v", got, s.want)
			}
		})
	}
}
