package graph

import (
	"testing"
)

func TestGetIdx(t *testing.T) {
	// Test Scenario 1: Target Found in the Middle of the Slice
	t.Log("Test Scenario 1: Target Found in the Middle of the Slice")
	target := 5
	nodes := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedIdx := 4

	idx := GetIdx(target, nodes)

	if idx != expectedIdx {
		t.Errorf("Expected index %d, got %d", expectedIdx, idx)
	}

	// Test Scenario 2: Target Found at the Beginning of the Slice
	t.Log("Test Scenario 2: Target Found at the Beginning of the Slice")
	target = 1
	nodes = []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedIdx = 0

	idx = GetIdx(target, nodes)

	if idx != expectedIdx {
		t.Errorf("Expected index %d, got %d", expectedIdx, idx)
	}

	// Test Scenario 3: Target Found at the End of the Slice
	t.Log("Test Scenario 3: Target Found at the End of the Slice")
	target = 8
	nodes = []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedIdx = 7

	idx = GetIdx(target, nodes)

	if idx != expectedIdx {
		t.Errorf("Expected index %d, got %d", expectedIdx, idx)
	}

	// Test Scenario 4: Target Not Found in the Slice
	t.Log("Test Scenario 4: Target Not Found in the Slice")
	target = 9
	nodes = []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedIdx = -1

	idx = GetIdx(target, nodes)

	if idx != expectedIdx {
		t.Errorf("Expected index %d, got %d", expectedIdx, idx)
	}

	// Test Scenario 5: Empty Slice
	t.Log("Test Scenario 5: Empty Slice")
	target = 1
	nodes = []int{}
	expectedIdx = -1

	idx = GetIdx(target, nodes)

	if idx != expectedIdx {
		t.Errorf("Expected index %d, got %d", expectedIdx, idx)
	}

	// Test Scenario 6: Target Value is Less Than the Minimum Value in the Slice
	t.Log("Test Scenario 6: Target Value is Less Than the Minimum Value in the Slice")
	target = 0
	nodes = []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedIdx = -1

	idx = GetIdx(target, nodes)

	if idx != expectedIdx {
		t.Errorf("Expected index %d, got %d", expectedIdx, idx)
	}

	// Test Scenario 7: Target Value is Greater Than the Maximum Value in the Slice
	t.Log("Test Scenario 7: Target Value is Greater Than the Maximum Value in the Slice")
	target = 9
	nodes = []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedIdx = -1

	idx = GetIdx(target, nodes)

	if idx != expectedIdx {
		t.Errorf("Expected index %d, got %d", expectedIdx, idx)
	}
}
