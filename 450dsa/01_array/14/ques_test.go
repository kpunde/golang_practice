package _14

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
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		},
		output: [][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		},
	},
	{
		input: [][]int{
			{1, 4},
			{4, 5},
		},
		output: [][]int{
			{1, 5},
		},
	},
}

func TestGetMerge(t *testing.T) {
	for index, test := range testCases {
		if op := getMerge(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
