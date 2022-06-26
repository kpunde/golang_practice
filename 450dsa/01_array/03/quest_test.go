package _3

import "testing"

type testCases struct {
	input    []int
	inputPos int
	output   int
}

var addTests = []testCases{
	{
		input:    []int{7, 10, 4, 3, 20, 15},
		inputPos: 3,
		output:   7,
	},
	{
		input:    []int{7, 10, 4, 20, 15},
		inputPos: 4,
		output:   15,
	},
}

func TestKSmallest(t *testing.T) {
	for _, test := range addTests {
		output := getKSmallest(test.input, test.inputPos)
		if output != test.output {
			t.Errorf("Test failed, expected %v, got %v", test.output, output)
		}
	}
}
