package _2

import "math"

func findMinMax(list []int) (int, int) {
	var min = math.MaxInt
	var max = math.MinInt
	for _, item := range list {
		if item < min {
			min = item
		}

		if item > max {
			max = item
		}
	}

	return min, max
}
