package _31

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{7},
			{2, 4, 3, 1, 2, 4, 3},
		},
		output: 2,
	},
	{
		input: [][]int{
			{7},
			{2, 3, 1, 2, 4, 3},
		},
		output: 2,
	},
	{
		input: [][]int{
			{4},
			{1, 4, 4},
		},
		output: 1,
	},
	{
		input: [][]int{
			{11},
			{1, 1, 1, 1, 1, 1, 1, 1},
		},
		output: 0,
	},
	{
		input: [][]int{
			{11},
			{1, 2, 3, 4, 5},
		},
		output: 3,
	},
	{
		input: [][]int{
			{7},
			{1, 1, 1, 1, 7},
		},
		output: 1,
	},
}

func TestMinSubArrayLen(t *testing.T) {
	for index, test := range testCases {
		if op := minSubArrayLenImproved(test.input[0][0], test.input[1]); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
