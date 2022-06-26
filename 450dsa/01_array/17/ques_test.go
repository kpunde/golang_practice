package _17

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  []int
	output []int
}

var testCases = []testCase{
	{
		input:  []int{1, 2, 3},
		output: []int{1, 3, 2},
	},
	{
		input:  []int{3, 2, 1},
		output: []int{1, 2, 3},
	},
	{
		input:  []int{1, 1, 5},
		output: []int{1, 5, 1},
	},
}

func TestNextPermutation(t *testing.T) {
	for index, test := range testCases {
		if op := nextPermutation(test.input); reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
