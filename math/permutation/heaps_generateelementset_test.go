package permutation_test

import (
	"fmt"
	"permutation"
	"testing"
)

// TestGenerateElementSet tests the GenerateElementSet function for various scenarios.
func TestGenerateElementSet(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		wantLen   int
		wantFirst string // Assuming we want to check the first element if the length is not zero
	}{
		{"PositiveN", 5, 5, "1"},
		{"ZeroN", 0, 0, ""},
		{"NegativeN", -1, 0, ""},
		// TODO: Adjust the value of 'n' to a number that could cause overflow if needed
		{"IntegerOverflowN", /* some very large number */, 0, ""},
		// Concurrency test is handled separately
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(chan []string)
			go permutation.GenerateElementSet(out, tt.n)
			got := <-out

			if len(got) != tt.wantLen {
				t.Errorf("GenerateElementSet() got len = %v, want len %v", len(got), tt.wantLen)
			} else if tt.wantLen > 0 && got[0] != tt.wantFirst {
				t.Errorf("GenerateElementSet() got first element = %v, want first element %v", got[0], tt.wantFirst)
			}
		})
	}

	// Concurrency test
	t.Run("Concurrency", func(t *testing.T) {
		n := 5
		concurrentCalls := 10
		channels := make([]chan []string, concurrentCalls)
		expectedFirst := "1"

		for i := range channels {
			channels[i] = make(chan []string)
			go permutation.GenerateElementSet(channels[i], n)
		}

		for _, ch := range channels {
			select {
			case got := <-ch:
				if len(got) != n {
					t.Errorf("GenerateElementSet() got len = %v, want len %v", len(got), n)
				} else if got[0] != expectedFirst {
					t.Errorf("GenerateElementSet() got first element = %v, want first element %v", got[0], expectedFirst)
				}
			default:
				t.Error("GenerateElementSet() did not send output on channel")
			}
		}
	})
}

// Note: Test for integer overflow is not included as it is not practical to test this in real scenarios.
// An actual integer overflow would depend on the system architecture and it's not feasible to generate a test case for it.
