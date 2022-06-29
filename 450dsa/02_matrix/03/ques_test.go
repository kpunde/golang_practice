package _03

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 3, 5},
			{2, 6, 9},
			{3, 6, 9},
		},
		output: 5,
	},
	{
		input: [][]int{
			{1},
			{2},
			{3},
		},
		output: 2,
	},
}

func TestGetMedian(t *testing.T) {
	for index, test := range testCases {
		if op := getMedian(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
