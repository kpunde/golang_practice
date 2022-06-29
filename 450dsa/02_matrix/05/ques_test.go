package _05

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
			{10, 20, 30, 40},
			{15, 25, 35, 45},
			{27, 29, 37, 48},
			{32, 33, 39, 50},
		},
		output: [][]int{
			{10, 15, 20, 25},
			{27, 29, 30, 32},
			{33, 35, 37, 39},
			{40, 45, 48, 50},
		},
	},
	{
		input: [][]int{
			{1, 5, 3},
			{2, 8, 7},
			{4, 6, 9},
		},
		output: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	},
}

func TestGetSortedMatrix(t *testing.T) {
	for index, test := range testCases {
		if op := getSortedMatrix(test.input); !reflect.DeepEqual(op, test.output) {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
