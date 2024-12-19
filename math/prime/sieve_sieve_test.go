package prime

import (
	"fmt"
	"os"
	"prime"
	"testing"
	"time"
)

func TestSieve(t *testing.T) {

	tests := []struct {
		name       string
		input      []int
		prime      int
		wantOutput []int
	}{
		{
			name:       "Filtering out non-prime numbers",
			input:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			prime:      2,
			wantOutput: []int{3, 5, 7, 9},
		},
		{
			name:       "Handling of prime numbers in the input channel",
			input:      []int{3, 5, 7, 11, 13, 17},
			prime:      2,
			wantOutput: []int{3, 5, 7, 11, 13, 17},
		},
		{
			name:       "Zero and negative numbers handling",
			input:      []int{-1, 0, 1, 2, 3, 4},
			prime:      2,
			wantOutput: []int{1, 3},
		},
		{
			name:       "Sieve function with a non-prime filtering number",
			input:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			prime:      4,
			wantOutput: []int{2, 3, 5, 7, 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := make(chan int, len(tc.input))
			out := make(chan int, len(tc.input))

			go prime.Sieve(in, out, tc.prime)

			for _, num := range tc.input {
				in <- num
			}
			close(in)

			var gotOutput []int
			for range tc.input {
				select {
				case result := <-out:
					gotOutput = append(gotOutput, result)
				case <-time.After(1 * time.Second):
					t.Fatal("Timeout: Sieve function took too long to process")
				}
			}

			for _, want := range tc.wantOutput {
				found := false
				for _, got := range gotOutput {
					if got == want {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Test '%s': expected %v to include %d, but it did not", tc.name, gotOutput, want)
				}
			}

			t.Logf("Test '%s': success. Expected output: %v, Got output: %v", tc.name, tc.wantOutput, gotOutput)
		})
	}
}

