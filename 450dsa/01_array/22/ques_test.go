package _22

import "testing"

type testCase struct {
	input  []int
	output bool
}

var testCases = []testCase{
	{
		input:  []int{4, 2, -3, 1, 6},
		output: true,
	},
	{
		input:  []int{2, 4, 0, 7, 5},
		output: true,
	},
	{
		input:  []int{2, 4, 7, 5},
		output: false,
	},
}

func TestSubArrayExists(t *testing.T) {
	for index, test := range testCases {
		if op := subArrayExists(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
