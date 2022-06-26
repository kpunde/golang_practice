package _29

import "testing"

type testCase struct {
	input  []int
	output int
}

var testCases = []testCase{
	{
		input:  []int{3, 0, 0, 2, 0, 4},
		output: 10,
	},
	{
		input:  []int{7, 4, 0, 9},
		output: 10,
	},
	{
		input:  []int{6, 9, 9},
		output: 0,
	},
}

func TestTrappingWater(t *testing.T) {
	for index, test := range testCases {
		if op := trappingWater(test.input); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
