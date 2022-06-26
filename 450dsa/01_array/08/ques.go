package _8

import "math"

// Kadane's Algo. Whenever we get a local sum as negative, discard that part and move forward
func getLargestSum(list []int) int {
	gMax := math.MinInt
	lMax := 0
	for _, item := range list {
		lMax += item
		gMax = int(math.Max(float64(lMax), float64(gMax)))
		if lMax < 0 {
			lMax = 0
		}
	}
	return gMax
}
