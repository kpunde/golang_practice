package _21

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
		input:  []int{9, 4, -2, -1, 5, 0, -5, -3, 2},
		output: []int{9, -2, 4, -1, 5, -5, 0, -3, 2},
	},
	{
		input:  []int{-5, -2, 5, 2, 4, 7, 1, 8, 0, -8},
		output: []int{5, -5, 2, -2, 4, -8, 7, 1, 8, 0},
	},
}

func TestGetRearrange(t *testing.T) {
	for index, test := range testCases {
		if op := getRearrange(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}