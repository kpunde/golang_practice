package _34

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{10, 999, 10},
		output: 0,
	},
	{
		input:  []int{10, 99, 44, 10},
		output: 1,
	},
}

func TestMerge(t *testing.T) {
	for index, test := range testCases {
		if op := merge(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
