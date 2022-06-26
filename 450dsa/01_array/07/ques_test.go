package _7

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  []int
	output []int
}

var testCases = []testCase{
	{
		input:  []int{1, 2, 3, 4, 5},
		output: []int{5, 1, 2, 3, 4},
	},

	{
		input:  []int{9, 8, 7, 6, 4, 2, 1, 3},
		output: []int{3, 9, 8, 7, 6, 4, 2, 1},
	},
}

func TestGetCyclicRotate(t *testing.T) {
	for _, test := range testCases {
		if output := getCyclicRotate(test.input); !reflect.DeepEqual(output, test.output) {
			t.Errorf("Output %v not equal to expected %v", output, test.output)
		}
	}
}
