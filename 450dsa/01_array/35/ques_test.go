package _35

import "testing"

type testCase struct {
	input  [][]int
	output float64
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 3},
			{2},
		},
		output: 2,
	},
	{
		input: [][]int{
			{1, 2},
			{3, 4},
		},
		output: 2.500,
	},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for index, test := range testCases {
		if op := findMedianSortedArrays(test.input[0], test.input[1]); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
