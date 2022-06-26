package _32

import (
	"fmt"
	"testing"
)

type testCase struct {
	input [][]int
}

var testCases = []testCase{
	{
		input: [][]int{
			{1, 2, 3, 3, 4},
			{1, 2},
		},
	},
	{
		input: [][]int{
			{1, 2, 3},
			{1, 3},
		},
	},
	{
		input: [][]int{
			{15, 9, 7, 5, 2, 8, 6, 14, 20},
			{7, 14},
		},
	},
}

func TestThreeWayPartition(t *testing.T) {
	for _, test := range testCases {
		fmt.Println(threeWayPartition(test.input[0], test.input[1][0], test.input[1][1]))
	}
}
