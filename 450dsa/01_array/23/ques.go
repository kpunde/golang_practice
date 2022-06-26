package _23

import "sort"

func getLongestSubsequence(arr []int) int {
	sort.Ints(arr)

	gLen := 0
	lLen := 0
	nextSq := arr[0]

	for _, item := range arr {
		if item == nextSq {
			lLen++
			nextSq++
		} else {
			lLen = 0
			nextSq = item + 1
		}

		if lLen > gLen {
			gLen = lLen
		}
	}

	return gLen
}
