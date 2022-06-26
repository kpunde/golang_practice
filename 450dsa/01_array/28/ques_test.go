package _28

import "testing"

type testCase struct {
	input  [][]int
	output bool
}

var testCases = []testCase{
	{
		input: [][]int{
			{13},
			{1, 4, 45, 6, 10, 8},
		},
		output: true,
	},
	{
		input: [][]int{
			{10},
			{1, 2, 4, 3, 6},
		},
		output: true,
	},
}

func TestFind3Numbers(t *testing.T) {
	for index, test := range testCases {
		if op := find3Numbers(test.input[1], test.input[0][0]); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
