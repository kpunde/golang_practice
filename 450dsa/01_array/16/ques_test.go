package _16

import (
	"testing"
)

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{2, 4, 1, 3, 5},
		output: 3,
	},
	{
		input:  []int{2, 3, 4, 5, 6},
		output: 0,
	},
	{
		input:  []int{1, 1, 1},
		output: 0,
	},
}

func TestInversionCount(t *testing.T) {
	for index, test := range testCases {
		if op := getInversionCount(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
