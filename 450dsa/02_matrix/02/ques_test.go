package _02

import (
	"testing"
)

type testCase struct {
	input  [][]int
	target int
	output bool
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		target: 3,
		output: true,
	},
	{
		input: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		target: 13,
		output: false,
	},
}

func TestGetSearchMatrix(t *testing.T) {
	for index, test := range testCases {
		if op := getSearchMatrix(test.input, test.target); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
