package _2

import (
	"testing"
)

type testResults struct {
	input, output []int
}

var addTests = []testResults{
	{
		input:  []int{1, 2, 4, 5},
		output: []int{1, 5},
	},
	{
		input:  []int{0, 0, 0, 0},
		output: []int{0, 0},
	},
	{
		input:  []int{1, 1, 1, 1},
		output: []int{1, 1},
	},
	{
		input:  []int{1, 1, 3, 3},
		output: []int{1, 3},
	},
	{
		input:  []int{100, 22, 43, 54},
		output: []int{22, 100},
	},
}

func TestReverseList(t *testing.T) {
	for _, test := range addTests {
		if min, max := findMinMax(test.input); min != test.output[0] || max != test.output[1] {
			t.Errorf("Test failed, expected min=%v, max=%v and got min=%v, max=%v", test.output[0], test.output[1], min, max)
		}
	}
}
