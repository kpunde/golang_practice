package _27

import "testing"

type testCase struct {
	input  [][]int
	output bool
}

var testCases = []testCase{
	{
		input: [][]int{
			{11, 1, 13, 21, 3, 7},
			{11, 3, 7, 1},
		},
		output: true,
	},
	{
		input: [][]int{
			{1, 2, 3, 4, 5, 6},
			{1, 2, 4},
		},
		output: true,
	},
	{
		input: [][]int{
			{10, 5, 2, 23, 19},
			{19, 5, 3},
		},
		output: false,
	},
}

func TestGetIsSubset(t *testing.T) {
	for index, test := range testCases {
		if op := getIsSubset(test.input[0], test.input[1]); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
