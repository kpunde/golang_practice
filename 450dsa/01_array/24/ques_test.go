package _24

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  int
	output []int
}

var testCases = []testCase{
	{
		input:  5,
		output: []int{1, 2, 0},
	},
	{
		input:  10,
		output: []int{3, 6, 2, 8, 8, 0, 0},
	},
	{
		input:  20,
		output: []int{2, 4, 3, 2, 9, 0, 2, 0, 0, 8, 1, 7, 6, 6, 4, 0, 0, 0, 0},
	},
}

func TestGetFactorial(t *testing.T) {
	for index, test := range testCases {
		if op := getFactorial(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
