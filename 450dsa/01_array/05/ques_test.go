package _5

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
		input:  []int{1, -1, 3, 2, -7, -5, 11, 6},
		output: []int{1, 3, 2, 11, 6, -1, -7, -5},
	},
	{
		input:  []int{-5, 7, -3, -4, 9, 10, -1, 11},
		output: []int{7, 9, 10, 11, -5, -3, -4, -1},
	},
}

func TestMoveNegativeToEnd(t *testing.T) {
	for _, test := range testCases {
		if op := moveNegativeToEnd(test.input); !reflect.DeepEqual(test.output, op) {
			t.Errorf("Output %v not equal to expected %v", op, test.output)
		}
	}
}
