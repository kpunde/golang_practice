package _13

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{1, 2, 3, -2, 5},
		output: 9,
	},
	{
		input:  []int{-1, -2, -3, -4},
		output: -1,
	},
	{
		input:  []int{-5, 4, 6, -3, 4, -1},
		output: 11,
	},
}

func TestGetMaxSubarraySum(t *testing.T) {
	for index, test := range testCases {
		if op := getMaxSubarraySum(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
