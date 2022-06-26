package _6

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 2, 3, 4, 5},
			{1, 2, 3},
		},
		output: 5,
	},
	{
		input: [][]int{
			{85, 25, 1, 32, 54, 6},
			{85, 2},
		},
		output: 7,
	},
}

func TestGetUnion(t *testing.T) {
	for _, test := range testCases {
		if output := getUnion(test.input[0], test.input[1]); output != test.output {
			t.Errorf("Output %v not equal to expected %v", output, test.output)
		}
	}
}
