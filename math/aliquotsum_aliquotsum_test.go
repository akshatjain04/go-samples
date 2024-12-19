package math

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

var ErrNonZeroArgsOnly = errors.New("non-zero arguments only")
var ErrPosArgsOnly = errors.New("positive arguments only")func TestAliquotSum(t *testing.T) {

	testCases := []struct {
		Name          string
		Input         int
		ExpectedSum   int
		ExpectedError error
	}{
		{
			Name:          "Positive input with expected aliquot sum",
			Input:         28,
			ExpectedSum:   28,
			ExpectedError: nil,
		},
		{
			Name:          "Negative input returns error",
			Input:         -1,
			ExpectedSum:   0,
			ExpectedError: ErrPosArgsOnly,
		},
		{
			Name:          "Zero input returns error",
			Input:         0,
			ExpectedSum:   0,
			ExpectedError: ErrNonZeroArgsOnly,
		},
		{
			Name:          "Large positive input",
			Input:         100,
			ExpectedSum:   117,
			ExpectedError: nil,
		},
		{
			Name:          "Input is a prime number",
			Input:         13,
			ExpectedSum:   1,
			ExpectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			sum, err := math.AliquotSum(tc.Input)

			if sum != tc.ExpectedSum || !errors.Is(err, tc.ExpectedError) {
				t.Errorf("AliquotSum(%d) = %d, %v; want %d, %v",
					tc.Input, sum, err, tc.ExpectedSum, tc.ExpectedError)
			}

			if err != nil {
				t.Logf("Received an error as expected for input %d: %v", tc.Input, err)
			} else {
				t.Logf("Received the correct aliquot sum for input %d: %d", tc.Input, sum)
			}
		})
	}
}
