package _1

import (
	"reflect"
	"testing"
)

type testResults struct {
	input, output []int
}

var addTests = []testResults{
	{
		input:  []int{1, 2, 4, 5},
		output: []int{5, 4, 2, 1},
	},
	{
		input:  []int{0, 0, 0, 0},
		output: []int{0, 0, 0, 0},
	},
	{
		input:  []int{1, 1, 1, 1},
		output: []int{1, 1, 1, 1},
	},
	{
		input:  []int{1, 1, 3, 3},
		output: []int{3, 3, 1, 1},
	},
	{
		input:  []int{11, 22, 43, 54},
		output: []int{54, 43, 22, 11},
	},
}

func TestReverseList(t *testing.T) {
	for _, test := range addTests {
		if op := reverseList(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Output %v not equal to expected %v", op, test.output)
		}
	}
}
