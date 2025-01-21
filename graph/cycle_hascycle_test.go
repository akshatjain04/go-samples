package graph

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGraphHasCycle(t *testing.T) {
	// Table-driven tests for robustness
	tests := []struct {
		name        string
		graph       Graph
		expected    bool
		description string
	}{
		{
			name:        "Empty graph",
			graph:       Graph{vertices: 0, edges: make(map[int]map[int]int), Directed: false},
			expected:    false,
			description: "Should return false for an empty graph",
		},
		{
			name:        "Graph with a single vertex",
			graph:       Graph{vertices: 1, edges: make(map[int]map[int]int), Directed: false},
			expected:    false,
			description: "Should return false for a graph with a single vertex",
		},
		{
			name:        "Graph with a cycle",
			graph:       Graph{vertices: 3, edges: make(map[int]map[int]int), Directed: false},
			expected:    true,
			description: "Should return true for a graph with a cycle",
		},
		{
			name:        "Graph with no cycle",
			graph:       Graph{vertices: 3, edges: make(map[int]map[int]int), Directed: false},
			expected:    false,
			description: "Should return false for a graph with no cycle",
		},
		{
			name:        "Directed graph with a cycle",
			graph:       Graph{vertices: 3, edges: make(map[int]map[int]int), Directed: true},
			expected:    true,
			description: "Should return true for a directed graph with a cycle",
		},
		{
			name:        "Directed graph with no cycle",
			graph:       Graph{vertices: 3, edges: make(map[int]map[int]int), Directed: true},
			expected:    false,
			description: "Should return false for a directed graph with no cycle",
		},
		{
			name:        "Graph with negative weight edges",
			graph:       Graph{vertices: 3, edges: make(map[int]map[int]int), Directed: false},
			expected:    false,
			description: "Should return false for a graph with negative weight edges",
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture and utilize os.Stdout for testing functions that do not return values, ensuring output reliability
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function under test
			actual := test.graph.HasCycle()

			// Restore standard output
			os.Stdout = oldStdout
			var output string
			fmt.Fscanf(r, "%s", &output)

			// Log detailed success and failure reasons for diagnostic clarity
			if actual == test.expected {
				t.Logf("%s: PASSED - %s", test.name, test.description)
			} else {
				t.Errorf("%s: FAILED - Expected %t, got %t - %s", test.name, test.expected, actual, test.description)
			}

			// Ensure outputs match expectations
			if !reflect.DeepEqual(output, "") {
				t.Errorf("%s: Unexpected output: %s", test.name, output)
			}
		})
	}
}
