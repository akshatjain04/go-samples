package prime

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {

	testCases := []struct {
		name     string
		limit    int
		expected []int
	}{
		{
			name:     "Zero Limit Scenario",
			limit:    0,
			expected: []int{},
		},
		{
			name:     "Negative Limit Scenario",
			limit:    -1,
			expected: []int{},
		},
		{
			name:     "Single Prime Scenario",
			limit:    1,
			expected: []int{2},
		},
		{
			name:     "Large Limit Scenario",
			limit:    10000,
			expected: nil,
		},
		{
			name:     "Consecutive Calls Consistency",
			limit:    100,
			expected: nil,
		},
		{
			name:     "Prime Generation Algorithm Correctness",
			limit:    100,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Generate(tc.limit)
			if tc.expected != nil {
				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("TestGenerate (%s): expected %v, got %v", tc.name, tc.expected, result)
				} else {
					t.Logf("TestGenerate (%s): passed", tc.name)
				}
			} else {
				if len(result) != tc.limit {
					t.Errorf("TestGenerate (%s): expected length %d, got length %d", tc.name, tc.limit, len(result))
				} else {
					t.Logf("TestGenerate (%s): passed", tc.name)
				}
			}
		})
	}
}
