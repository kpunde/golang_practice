package _18

import "math"

func getMaxProfit(prices []int) int {
	minima, maxDiff := prices[0], 0
	for _, item := range prices {
		minima = int(math.Min(float64(minima), float64(item)))
		maxDiff = int(math.Max(float64(maxDiff), float64(item-minima)))
	}
	return maxDiff
}
