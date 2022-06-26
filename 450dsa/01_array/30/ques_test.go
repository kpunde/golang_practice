package _30

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{5},
			{3, 4, 1, 9, 56, 7, 9, 12},
		},
		output: 6,
	},
	{
		input: [][]int{
			{3},
			{7, 3, 2, 4, 9, 12, 56},
		},
		output: 2,
	},
}

func TestTrappingWater(t *testing.T) {
	for index, test := range testCases {
		if op := findMinDiff(test.input[1], test.input[0][0]); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
