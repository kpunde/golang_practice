package _23

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{2, 6, 1, 9, 4, 5, 3},
		output: 6,
	},
	{
		input:  []int{1, 9, 3, 10, 4, 20, 2},
		output: 4,
	},
}

func TestGetLongestSubsequence(t *testing.T) {
	for _, test := range testCases {
		if op := getLongestSubsequence(test.input); op != test.output {
			t.Errorf("Output %v not equal to expected %v", op, test.output)
		}
	}
}
