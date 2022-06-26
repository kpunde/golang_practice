package _0

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9},
		output: 3,
	},
	{
		input:  []int{1, 4, 3, 2, 6, 7},
		output: 2,
	},
	{
		input:  []int{1, 3, 3, 2, 0, 0},
		output: -1,
	},
}

func TestGetMinJumps(t *testing.T) {
	for _, test := range testCases {
		if op := getMinJumps(test.input); op != test.output {
			t.Errorf("Output %v not equal to expected %v", op, test.output)
		}
	}
}
