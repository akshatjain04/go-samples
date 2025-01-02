package sort

import (
	"testing"
	"reflect"
)

// TestCount tests the Count sorting method
func TestCount(t *testing.T) {
	type test struct {
		input  []int
		output []int
	}

	tests := []test{
		{input: []int{}, output: []int{}},
		{input: []int{3}, output: []int{3}},
		{input: []int{6, 6, 6}, output: []int{6, 6, 6}},
		{input: []int{4, 3, 2}, output: []int{2, 3, 4}},
		{input: []int{-1, -3, -2}, output: []int{-3, -2, -1}},
	}

	for i, scenario := range tests {
		result := Count(scenario.input)

		if !reflect.DeepEqual(result, scenario.output) {
			t.Errorf("Test Scenario %d failed - Results not match\nExpected output: %v\nObtained output: %v", i+1, scenario.output, result)
		} else {
			t.Logf("Test Scenario %d passed", i+1)
		}
	}
}
