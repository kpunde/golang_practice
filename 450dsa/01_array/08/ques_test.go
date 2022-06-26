package _8

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{1, 2, 3, -2, 5},
		output: 9,
	},
	{
		input:  []int{-1, -2, -3, -4},
		output: -1,
	},
	{
		input:  []int{-5, 4, 6, -3, 4, -1},
		output: 11,
	},
}

func TestGetLargestSum(t *testing.T) {
	for _, test := range testCases {
		if op := getLargestSum(test.input); op != test.output {
			t.Errorf("Output %v not equal to expected %v", op, test.output)
		}
	}
}
