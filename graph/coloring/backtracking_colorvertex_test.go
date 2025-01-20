package coloring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphColorVertex(t *testing.T) {
	// Scenario 1: Coloring a graph with no vertices
	t.Run("no vertices", func(t *testing.T) {
		// Arrange
		graph := &Graph{vertices: 0}

		// Act
		result := graph.colorVertex(0, make(map[int]Color))

		// Assert
		assert.False(t, result, "Expected false for a graph with no vertices")
	})

	// Scenario 2: Coloring a graph with one vertex
	t.Run("one vertex", func(t *testing.T) {
		// Arrange
		graph := &Graph{vertices: 1}

		// Act
		result := graph.colorVertex(0, make(map[int]Color))

		// Assert
		assert.True(t, result, "Expected true for a graph with one vertex")
	})

	// Scenario 3: Coloring a graph with multiple vertices and no edges
	t.Run("multiple vertices and no edges", func(t *testing.T) {
		// Arrange
		graph := &Graph{vertices: 5}

		// Act
		result := graph.colorVertex(0, make(map[int]Color))

		// Assert
		assert.True(t, result, "Expected true for a graph with multiple vertices and no edges")
	})

	// Scenario 4: Coloring a graph with a cycle
	t.Run("cycle", func(t *testing.T) {
		// Arrange
		graph := &Graph{
			vertices: 4,
			edges: map[int]map[int]struct{}{
				0: {1: {}},
				1: {2: {}},
				2: {3: {}},
				3: {0: {}},
			},
		}

		// Act
		result := graph.colorVertex(0, make(map[int]Color))

		// Assert
		assert.True(t, result, "Expected true for a graph with a cycle")
	})

	// Scenario 5: Coloring a graph with insufficient colors
	t.Run("insufficient colors", func(t *testing.T) {
		// Arrange
		graph := &Graph{vertices: 5}

		// Act
		result := graph.colorVertex(0, map[int]Color{0: 1, 1: 2, 2: 3})

		// Assert
		assert.False(t, result, "Expected false for a graph with insufficient colors")
	})

	// Scenario 6: Coloring a graph with an invalid color
	t.Run("invalid color", func(t *testing.T) {
		// Arrange
		graph := &Graph{vertices: 3}

		// Assert
		assert.Panics(t, func() { graph.colorVertex(0, map[int]Color{0: 1, 1: 2, 2: 4}) }, "Expected panic for an invalid color")
	})
}
