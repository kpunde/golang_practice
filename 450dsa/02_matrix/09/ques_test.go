package _09

import (
	"testing"
)

type testCase struct {
	input  [][]int
	k      int
	output int
}

var testCases = []testCase{
	{
		input: [][]int{
			{16, 28, 60, 64},
			{22, 41, 63, 91},
			{27, 50, 87, 93},
			{36, 78, 87, 94},
		},
		k:      3,
		output: 27,
	},
	{
		input: [][]int{
			{10, 20, 30, 40},
			{15, 25, 35, 45},
			{24, 29, 37, 48},
			{32, 33, 39, 50},
		},
		k:      7,
		output: 30,
	},
}

func TestGetKthSmallest(t *testing.T) {
	for index, test := range testCases {
		if op := getKthSmallest(test.input, test.k); op != test.output {
			t.Errorf("Error at test index: %v Output %v not equal to expected %v", index, op, test.output)
		} else {
			t.Logf("Index %v passed successfully", index)
		}
	}
}
