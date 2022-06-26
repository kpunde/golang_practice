package _19

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{6},
			{1, 5, 7, 1},
		},
		output: 2,
	},
	{
		input: [][]int{
			{2},
			{1, 1, 1, 1},
		},
		output: 6,
	},
}

func TestGetPairsCount(t *testing.T) {
	for index, test := range testCases {
		if op := getPairsCount(test.input[1], test.input[0][0]); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
