package _4

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
		input:  []int{0, 1, 0, 1, 0, 2, 2, 2, 0, 0, 0, 2, 2, 1, 1, 1},
		output: []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
	},
}

func TestSort012(t *testing.T) {
	for _, test := range testCases {
		if result := sort012(test.input); !reflect.DeepEqual(result, test.output) {
			t.Errorf("Output %v not equal to expected %v", result, test.output)
		}
	}
}
