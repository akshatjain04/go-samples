// ********RoostGPT********
/*
Test generated by RoostGPT for test golang-test-1 using AI Type  and AI Model 

ROOST_METHOD_HASH=least_a3eb2b15d7
ROOST_METHOD_SIG_HASH=least_288573aa58

Scenario 1: Test with both input slices q1 and q2 empty

Details:
  Description: This test aims to checks how the function behaves when both input slices q1 and q2 are empty. Since there's no node to return, the outcome of the function is undefined.
Execution:
  Arrange: Create two empty slices of Nodes.
  Act: Invoke 'least' function giving these two slices as parameters.
  Assert: Using the Go testing facilities, assert that the panic or error is thrown as expected.
Validation:
  The assertion is chosen because we want to validate if function handles the zero length array inputs. This is achievable by catching the panic error or any error that function throws. This edge case is important, because it handles scenarios where no nodes are given to parse.

Scenario 2: Test with non-empty q1 and empty q2

Details:
  Description: The test aims to validate the functionality of the 'least' function when the slice q2 is empty but q1 is not. 
Execution:
  Arrange: Create a non-empty slice q1 with Nodes and an empty slice q2.
  Act: Invoke 'least' function with these two slices as parameters.
  Assert: Assert that the returned Node is the first Node of q1, the remaining slice of q1 is q1 with the first Node excluded, and the returned slice for q2 is same as the empty slice q2.
Validation:
  We choose this assertion to validate that the function performs as expected when one of the slices is empty. This checks that the function returns the first Node of the non-empty slice along with the updated slices. This is crucial for ensuring the proper operation of the "least" function in scenarios where there is no available node from one of the queues.

Scenario 3: Test with both input slices q1 and q2 non-empty and q1's first node's weight is less or equal to q2's first node's weight

Details:
  Description: The test aims to validate the functionality for the normal operation where both input slices are non-empty and q1's first node's weight is less or equal to that of q2.
Execution:
  Arrange: Initialize two slices q1 and q2 with Nodes such that q1's first node's weight is less or equal to q2's first node's weight.
  Act: Call 'least' function with the two slices as parameters.
  Assert: Assert that the Node returned is q1's first Node, q1 without the first Node, and q2 remains the same.
Validation:
  This assertion is chosen to validate that the function behaves as intended in the scenario where q1's head node has a lesser or equal weight to q2's head node. This is important as it validates the core logic of the function where it's supposed to return the least weighted node. 

Scenario 4: Test with both input slices q1 q2 non-empty and q1's first node weight is greater than q2's first node weight

Details:
  Description: This scenario is meant to validate the normal operation involving the case where both input slices are non-empty and q1's first Node's weight is greater than that of q2.
Execution:
  Arrange: Initialize slices q1 and q2 with Nodes such that q1's first Node's weight is greater than q2's first Node's weight.
  Act: Call 'least' function with these two slices as parameters.
  Assert: Assert that Node returned is the first Node from q2, the remaining slice of q1 remains the same and the remaining slice for q2 is q2 excluding the returned first Node.
Validation:
  This assertion is crucial as it verifies the correct selection of the least weighted node from two non-empty queues. This is an important functionality of the "least" function - it ensures that the function will always return the node with the lowest weight when presented with two non-empty queues.

*/

// ********RoostGPT********
package compression

import (
	"testing"
)

func TestLeast(t *testing.T) {
	tcs := []struct {
		name     string
		q1       []Node
		q2       []Node
		expected Node
		expQ1    []Node
		expQ2    []Node
	}{
		{
			name:     "BothQueuesEmpty",
			q1:       []Node{},
			q2:       []Node{},
			expected: Node{},
			expQ1:    []Node{},
			expQ2:    []Node{},
		},
		{
			name:     "FirstNonEmptySecondEmpty",
			q1:       []Node{{symbol: 'a', weight: 5}},
			q2:       []Node{},
			expected: Node{symbol: 'a', weight: 5},
			expQ1:    []Node{},
			expQ2:    []Node{},
		},
		{
			name: "BothNonEmptyFirstLessOrEqual",
			q1:   []Node{{symbol: 'a', weight: 5}, {symbol: 'b', weight: 6}},
			q2:   []Node{{symbol: 'c', weight: 7}},
			expected: Node{symbol: 'a', weight: 5},
			expQ1:    []Node{{symbol: 'b', weight: 6}},
			expQ2:    []Node{{symbol: 'c', weight: 7}},
		},
		{
			name: "BothNonEmptyFirstGreater",
			q1:   []Node{{symbol: 'd', weight: 8}},
			q2:   []Node{{symbol: 'e', weight: 7}, {symbol: 'f', weight: 9}},
			expected: Node{symbol: 'e', weight: 7},
			expQ1:    []Node{{symbol: 'd', weight: 8}},
			expQ2:    []Node{{symbol: 'f', weight: 9}},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			gotNode, gotQ1, gotQ2 := least(tc.q1, tc.q2)

			if gotNode != tc.expected {
				t.Fatalf("Expected node: %+v, Got: %+v", tc.expected, gotNode)
			}
			if !equal(gotQ1, tc.expQ1) {
				t.Fatalf("Expected Q1: %+v, Got: %+v", tc.expQ1, gotQ1)
			}
			if !equal(gotQ2, tc.expQ2) {
				t.Fatalf("Expected Q2: %+v, Got: %+v", tc.expQ2, gotQ2)
			}
		})
	}

}

func equal(a, b []Node) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
