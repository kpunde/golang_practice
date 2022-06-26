package _9

import (
	"math"
	"sort"
)

func getMinimizedHeight(list []int, k int) int {
	sort.Ints(list)

	var min = 0
	var max = 0

	diff := list[len(list)-1] - list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > k {
			max = int(math.Max(float64(list[i-1]+k), float64(list[len(list)-1]-k)))
			min = int(math.Min(float64(list[i]-k), float64(list[0]+k)))
			diff = int(math.Min(float64(diff), float64(max-min)))
		}
	}

	return diff
}
