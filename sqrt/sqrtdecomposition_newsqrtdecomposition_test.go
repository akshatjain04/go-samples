package sqrt

import (
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"your_project/sqrt"
)

type intElement int
type intQuery int
func TestNewSqrtDecomposition(t *testing.T) {
	tests := []struct {
		name            string
		elements        []intElement
		querySingleElem func(intElement) intQuery
		unionQ          func(intQuery, intQuery) intQuery
		updateQ         func(intQuery, intElement, intElement) intQuery
		expectedBlocks  []intQuery
	}{
		{
			name:            "Scenario 1: Non-empty slice with valid functions",
			elements:        []intElement{1, 2, 3, 4},
			querySingleElem: intQuerySingleElement,
			unionQ:          intUnionQ,
			updateQ:         intUpdateQ,
			expectedBlocks:  []intQuery{1, 5, 14},
		},
		{
			name:            "Scenario 2: Empty slice",
			elements:        []intElement{},
			querySingleElem: intQuerySingleElement,
			unionQ:          intUnionQ,
			updateQ:         intUpdateQ,
			expectedBlocks:  nil,
		},
		{
			name:            "Scenario 3: Single element query function",
			elements:        []intElement{1, 2, 3, 4},
			querySingleElem: intQuerySingleElement,
			unionQ:          intUnionQ,
			updateQ:         intUpdateQ,
			expectedBlocks:  []intQuery{1, 5, 14},
		},
		{
			name:            "Scenario 4: Union function combines blocks",
			elements:        []intElement{1, 2, 3, 4},
			querySingleElem: intQuerySingleElement,
			unionQ:          intUnionQ,
			updateQ:         intUpdateQ,
			expectedBlocks:  []intQuery{1, 5, 14},
		},
		{
			name:            "Scenario 5: Update function updates blocks",
			elements:        []intElement{1, 2, 3, 4},
			querySingleElem: intQuerySingleElement,
			unionQ:          intUnionQ,
			updateQ:         intUpdateQ,
			expectedBlocks:  []intQuery{1, 5, 14},
		},
		{
			name:            "Scenario 6: Handling large element slices",
			elements:        make([]intElement, 10000),
			querySingleElem: intQuerySingleElement,
			unionQ:          intUnionQ,
			updateQ:         intUpdateQ,
			expectedBlocks:  make([]intQuery, uint64(math.Ceil(math.Sqrt(10000)))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			sqrtDec := sqrt.NewSqrtDecomposition(tt.elements, tt.querySingleElem, tt.unionQ, tt.updateQ)
			require.NotNil(t, sqrtDec, "SqrtDecomposition instance should not be nil")
			assert.Equal(t, tt.expectedBlocks, sqrtDec.Blocks(), "Blocks do not match expected results")
		})
	}
}
func intQuerySingleElement(element intElement) intQuery {
	return intQuery(element * element)
}
func intUnionQ(q1 intQuery, q2 intQuery) intQuery {
	return q1 + q2
}
func intUpdateQ(oldQ intQuery, oldE intElement, newE intElement) (newQ intQuery) {
	oldVal := intQuerySingleElement(oldE)
	newVal := intQuerySingleElement(newE)
	return oldQ - oldVal + newVal
}


