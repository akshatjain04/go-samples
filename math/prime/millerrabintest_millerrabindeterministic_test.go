package prime

import (
	"fmt"
	"os"
	"testing"
	"github.com/TheAlgorithms/Go/math/modular"
	"github.com/TheAlgorithms/Go/math/prime"
)

func TestMillerRabinDeterministic(t *testing.T) {

	testCases := []struct {
		name          string
		input         int64
		expectedPrime bool
		expectError   bool
	}{
		{"Small Prime", 5, true, false},
		{"Small Non-prime", 4, false, false},
		{"Negative Number", -7, false, true},
		{"Zero", 0, false, true},
		{"Number One", 1, false, false},
		{"Large Prime", 2147483647, true, false},
		{"Large Non-prime", 2147483648, false, false},
		{"Possible Overflow", 9223372036854775807, false, false},
	}

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			isPrime, err := prime.MillerRabinDeterministic(tc.input)

			w.Close()

			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err)
			}

			if tc.expectError {
				t.Log("Received expected error for input:", tc.input)
				return
			}

			if isPrime != tc.expectedPrime {
				t.Errorf("Expected prime result: %v, got: %v for input: %v", tc.expectedPrime, isPrime, tc.input)
			} else {
				t.Logf("Correctly identified %v as prime: %v", tc.input, tc.expectedPrime)
			}

			var buf []byte
			if _, err := r.Read(buf); err != nil {
				t.Error("Failed to read from redirected stdout:", err)
			}
			output := string(buf)
			if output != "" {
				t.Log("Captured output:", output)
			}
		})

		r, w, _ = os.Pipe()
		os.Stdout = w
	}
}
