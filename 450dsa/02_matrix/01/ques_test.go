package _01

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  [][]int
	output []int
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		output: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
	},

	{
		input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		output: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
}

func TestGetSpiralTransversal(t *testing.T) {
	for index, test := range testCases {
		if op := getSpiralTransversal(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
