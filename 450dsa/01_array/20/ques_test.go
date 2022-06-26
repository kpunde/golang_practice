package _20

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
			{1, 5, 10, 20, 40, 80},
			{6, 7, 20, 80, 100},
			{3, 4, 15, 20, 30, 70, 80, 120},
		},
		output: []int{20, 80},
	},
	{
		input: [][]int{
			{1, 20, 30, 40, 900},
			{10, 20, 300, 800},
			{20, 20, 20, 20},
		},
		output: []int{20},
	},
}

func TestGetCommonElements(t *testing.T) {
	for index, test := range testCases {
		if op := getCommonElements(test.input[0], test.input[1], test.input[2]); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
