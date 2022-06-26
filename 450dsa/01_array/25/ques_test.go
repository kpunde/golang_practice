package _25

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{6, -3, -10, -200, 0, 2},
		output: 2000,
	},
	{
		input:  []int{2, 3, 4, 5, -1, 0},
		output: 120,
	},
}

func TestGetMaxProduct(t *testing.T) {
	for index, test := range testCases {
		if op := getMaxProduct(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
