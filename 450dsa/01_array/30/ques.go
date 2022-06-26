package _30

import (
	"math"
	"sort"
)

func findMinDiff(arr []int, m int) int {
	sort.Ints(arr)
	i, j := 0, m-1
	diff := math.MaxInt

	for ; j < len(arr); i, j = i+1, j+1 {
		diff = int(math.Min(float64(diff), float64(arr[j]-arr[i])))
	}

	return diff
}
