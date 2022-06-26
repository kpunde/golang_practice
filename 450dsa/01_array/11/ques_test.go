package _11

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{1, 3, 4, 2, 2},
		output: 2,
	},
	{
		input:  []int{3, 1, 3, 4, 2},
		output: 3,
	},
	{
		input:  []int{1, 1},
		output: 1,
	},
}

func TestFindDuplicateN1(t *testing.T) {
	for _, test := range testCases {
		if op := findDuplicateN1(test.input); op != test.output {
			t.Errorf("Output %v not equal to expected %v", op, test.output)
		}
	}
}
