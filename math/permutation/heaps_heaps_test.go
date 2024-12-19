package permutation_test

import (
	"fmt"
	"math"
	"permutation"
	"testing"
	"time"
)

// TestHeaps function will test the Heaps function for various scenarios.
func TestHeaps(t *testing.T) {
	// Table-driven tests to cover the scenarios
	tests := []struct {
		name        string
		n           int
		expectError bool
		expectedLen int
	}{
		{
			name:        "Zero elements permutation",
			n:           0,
			expectError: false,
			expectedLen: 0,
		},
		{
			name:        "Single element permutation",
			n:           1,
			expectError: false,
			expectedLen: 1,
		},
		{
			name:        "Negative elements permutation",
			n:           -1,
			expectError: true,
			expectedLen: 0,
		},
		// TODO: Update the expectedLen for "Large number of elements permutation" after confirming the factorial function's implementation
		{
			name:        "Large number of elements permutation",
			n:           10,
			expectError: false,
			expectedLen: int(math.Factorial(10)), // Assuming Factorial function is available and returns int
		},
		// TODO: Add more test cases for "Concurrent execution of Heaps function" if needed
		{
			name:        "Concurrent execution of Heaps function",
			n:           3,
			expectError: false,
			expectedLen: 6, // 3! permutations
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := make(chan []string)
			go permutation.Heaps(out, tc.n)

			select {
			case permutations := <-out:
				if tc.expectError && len(permutations) != 0 {
					t.Errorf("TestHeaps failed for %v: expected error or empty slice, got %v", tc.name, permutations)
				}
				if !tc.expectError && len(permutations) != tc.expectedLen {
					t.Errorf("TestHeaps failed for %v: expected %d permutations, got %d", tc.name, tc.expectedLen, len(permutations))
				}
			case <-time.After(1 * time.Second):
				if !tc.expectError {
					t.Errorf("TestHeaps failed for %v: timeout exceeded while waiting for permutations", tc.name)
				}
			}
		})
	}

	// Additional test for concurrent execution
	t.Run("Concurrent execution of Heaps function", func(t *testing.T) {
		concurrentRuns := 5
		out := make(chan []string, concurrentRuns)
		for i := 0; i < concurrentRuns; i++ {
			go permutation.Heaps(out, i+1)
		}

		for i := 0; i < concurrentRuns; i++ {
			select {
			case permutations := <-out:
				if len(permutations) != int(math.Factorial(i+1)) {
					t.Errorf("Concurrent TestHeaps failed: expected %d permutations for n=%d, got %d", int(math.Factorial(i+1)), i+1, len(permutations))
				}
			case <-time.After(1 * time.Second):
				t.Errorf("Concurrent TestHeaps failed: timeout exceeded while waiting for permutations")
			}
		}
	})
}
