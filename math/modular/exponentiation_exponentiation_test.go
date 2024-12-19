package modular

import (
	"fmt"
	"github.com/TheAlgorithms/Go/math/modular"
	"testing"
)

func TestExponentiation(t *testing.T) {

	testCases := []struct {
		name           string
		base           int64
		exponent       int64
		mod            int64
		expectedResult int64
		expectedError  error
	}{
		{
			name:           "Base case with positive exponent and non-trivial modulus",
			base:           2,
			exponent:       10,
			mod:            1000,
			expectedResult: 24,
			expectedError:  nil,
		},
		{
			name:           "Modulus equals 1",
			base:           5,
			exponent:       3,
			mod:            1,
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name:           "Negative exponent provided",
			base:           5,
			exponent:       -3,
			mod:            10,
			expectedResult: -1,
			expectedError:  modular.ErrorNegativeExponent,
		},
		{
			name:           "Integer overflow during calculation",
			base:           math.MaxInt64,
			exponent:       2,
			mod:            math.MaxInt64,
			expectedResult: -1,
			expectedError:  modular.ErrorIntOverflow,
		},
		{
			name:           "Zero base with positive exponent",
			base:           0,
			exponent:       5,
			mod:            10,
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name:           "Positive base with zero exponent",
			base:           7,
			exponent:       0,
			mod:            10,
			expectedResult: 1,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := modular.Exponentiation(tc.base, tc.exponent, tc.mod)

			t.Logf("Testing %s", tc.name)
			t.Logf("Base: %d, Exponent: %d, Modulus: %d", tc.base, tc.exponent, tc.mod)
			t.Logf("Expected result: %d, Expected error: %v", tc.expectedResult, tc.expectedError)
			t.Logf("Actual result: %d, Actual error: %v", result, err)

			if result != tc.expectedResult {
				t.Errorf("Expected result %d, got %d", tc.expectedResult, result)
			}
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
