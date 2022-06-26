package _26

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
		input:  []int{3, 2, 3, 3, 2, 1, 2},
		output: []int{3, 2},
	},
	{
		input:  []int{1},
		output: []int{1},
	},
	{
		input:  []int{1, 2},
		output: []int{1, 2},
	},
	{
		input:  []int{1, 1, 1, 3, 3, 2, 2, 2},
		output: []int{1, 2},
	},
}

func TestGetMajorityElement(t *testing.T) {
	for index, test := range testCases {
		if op := getMajorityElement(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}

	t.Log("Running Improved version ...")

	for index, test := range testCases {
		if op := getMajorityElementImproved(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
