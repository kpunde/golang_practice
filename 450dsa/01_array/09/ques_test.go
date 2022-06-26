package _9

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 5, 8, 10},
			{2},
		},
		output: 5,
	},
	{
		input: [][]int{
			{3, 9, 12, 16, 20},
			{3},
		},
		output: 11,
	},
	{
		input: [][]int{
			{2, 6, 3, 4, 7, 2, 10, 3, 2, 1},
			{5},
		},
		output: 7,
	},
}

func TestGetMinimizedHeight(t *testing.T) {
	for _, test := range testCases {
		if op := getMinimizedHeight(test.input[0], test.input[1][0]); op != test.output {
			t.Errorf("Output %v not equal to expected %v", op, test.output)
		}
	}
}
