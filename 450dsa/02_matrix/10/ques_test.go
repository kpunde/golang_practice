package _10

import (
	"fmt"
	"testing"
)

func TestPrintCommonElements(t *testing.T) {
	tp := [][]int{
		{1, 2, 1, 4, 8},
		{3, 7, 8, 5, 1},
		{8, 7, 7, 3, 1},
		{8, 1, 2, 7, 9},
	}

	fmt.Println(printCommonElements(tp))
}
