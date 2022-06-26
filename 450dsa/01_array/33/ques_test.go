package _33

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{2, 1, 5, 6, 3},
			{3},
		},
		output: 1,
	},
	{
		input: [][]int{
			{2, 7, 9, 5, 8, 7, 4},
			{6},
		},
		output: 2,
	},
}

func TestMinSwap(t *testing.T) {
	for index, test := range testCases {
		if op := minSwap(test.input[0], test.input[1][0]); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
