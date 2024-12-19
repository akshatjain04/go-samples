package math_test

import (
	"math"
	"testing"
)

func TestCombinations(t *testing.T) {
	testCases := []struct {
		name           string
		n, k           int
		expectedResult int
		expectedError  error
	}{
		{
			name:           "Valid combinations with n greater than k",
			n:              5,
			k:              3,
			expectedResult: 10,
			expectedError:  nil,
		},
		{
			name:           "Valid combinations with n equal to k",
			n:              4,
			k:              4,
			expectedResult: 1,
			expectedError:  nil,
		},
		{
			name:           "Negative n argument",
			n:              -1,
			k:              3,
			expectedResult: -1,
			expectedError:  math.ErrPosArgsOnly,
		},
		{
			name:           "Negative k argument",
			n:              5,
			k:              -1,
			expectedResult: -1,
			expectedError:  math.ErrPosArgsOnly,
		},
		{
			name:           "k greater than n",
			n:              3,
			k:              5,
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name:           "Zero combinations when k is 0",
			n:              7,
			k:              0,
			expectedResult: 1,
			expectedError:  nil,
		},
		{
			name:           "Large values of n and k",
			n:              100,
			k:              50,
			expectedResult: 538992043,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Testing n: %d, k: %d", tc.n, tc.k)
			result, err := math.Combinations(tc.n, tc.k)
			if result != tc.expectedResult {
				t.Errorf("Expected result %d, got %d", tc.expectedResult, result)
			}
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
			if err == nil {
				t.Logf("Passed: %s", tc.name)
			} else {
				t.Logf("Passed with expected error: %s", tc.name)
			}
		})
	}
}
