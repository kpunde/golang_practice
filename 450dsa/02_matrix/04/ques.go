package _04

import "math"

func rowWithMax1s(arr [][]int) int {
	var max = math.MinInt
	var indexmax = 0

	for i := 0; i < len(arr); i++ {
		cnt := 0

		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 1 {
				cnt = len(arr[i]) - j
				break
			}
		}

		if cnt > max {
			max = cnt
			indexmax = i
		}
	}

	return indexmax
}
