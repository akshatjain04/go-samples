package graph

import (
	"testing"
)

// TestNew verifies that the "New" function correctly creates a new graph with the specified number of vertices.
func TestNew(t *testing.T) {
	// Create a table-driven test to validate various scenarios.
	tests := []struct {
		name     string
		vertices int
		expected *Graph
	}{
		{
			name:     "Create graph with valid vertices",
			vertices: 5,
			expected: &Graph{
				vertices: 5,
				edges:    make(map[int]map[int]int),
				Directed: false,
			},
		},
		{
			name:     "Create graph with zero vertices",
			vertices: 0,
			expected: &Graph{
				vertices: 0,
				edges:    make(map[int]map[int]int),
				Directed: false,
			},
		},
	}

	// Run the tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the "New" function with "v" as the argument.
			actual := New(tc.vertices)

			// Verify that the "vertices" field of "g" is equal to the value of "v".
			if actual.vertices != tc.expected.vertices {
				t.Errorf("Expected vertices: %d, got: %d", tc.expected.vertices, actual.vertices)
			}

			// Verify that the "edges" field of "g" is initialized correctly.
			if len(actual.edges) != 0 {
				t.Errorf("Expected edges to be empty, got: %v", actual.edges)
			}

			// Verify that the "Directed" field of "g" is set to false by default.
			if actual.Directed {
				t.Errorf("Expected Directed to be false, got: true")
			}
		})
	}
}
