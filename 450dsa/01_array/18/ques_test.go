package _18

import (
	"testing"
)

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{7, 1, 5, 3, 6, 4},
		output: 5,
	},
	{
		input:  []int{7, 6, 4, 3, 1},
		output: 0,
	},
}

func TestGetMaxProfit(t *testing.T) {
	for index, test := range testCases {
		if op := getMaxProfit(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
