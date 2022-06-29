package _08

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  [][]int
	output [][]int
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		output: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		},
	},
}

func TestGetRotate90Clockwise(t *testing.T) {
	for index, test := range testCases {
		if op := getRotate90Clockwise(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
