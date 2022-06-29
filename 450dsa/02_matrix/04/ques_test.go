package _04

import "testing"

type testCase struct {
	input  [][]int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
		},
		output: 2,
	},
	{
		input: [][]int{
			{0, 0},
			{1, 1},
		},
		output: 1,
	},
}

func TestRowWithMax1s(t *testing.T) {
	for index, test := range testCases {
		if op := rowWithMax1s(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
