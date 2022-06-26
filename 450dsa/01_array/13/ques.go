package _13

import "math"

func getMaxSubarraySum(arr []int) int {
	gSum, lSum := math.MinInt, 0
	for _, item := range arr {
		lSum += item
		gSum = int(math.Max(float64(gSum), float64(lSum)))
		if lSum < 0 {
			lSum = 0
		}
	}

	return gSum
}
