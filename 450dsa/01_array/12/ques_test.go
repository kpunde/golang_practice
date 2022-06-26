package _12

import (
	"reflect"
	"testing"
)

type testCase struct {
	input  [][]int
	output [][]int
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 3, 5, 7},
			{0, 2, 6, 8, 9},
		},
		output: [][]int{
			{0, 1, 2, 3},
			{5, 6, 7, 8, 9},
		},
	},
	{
		input: [][]int{
			{10, 12},
			{5, 18, 20},
		},
		output: [][]int{
			{5, 10},
			{12, 18, 20},
		},
	},
}

func TestMergeWithoutSpace(t *testing.T) {
	for index, test := range testCases {
		arr1, arr2 := mergeWithoutSpace(test.input[0], test.input[1])

		if !(reflect.DeepEqual(arr1, test.output[0]) && reflect.DeepEqual(arr2, test.output[1])) {
			t.Errorf("Failed at test index: %v, Output %v, %v not equal to expected %v, %v", index, arr1, arr2, test.output[0], test.output[1])
		} else {
			t.Logf("Test index %v passed successfully", index)
		}
	}
}
