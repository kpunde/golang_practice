package _25

import "math"

// Kadane algo extension

func getMaxProduct(arr []int) int {
	ans := arr[0]
	min, max := ans, ans

	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			max, min = min, max
		}

		max = int(math.Max(float64(arr[i]), float64(arr[i]*max)))
		min = int(math.Min(float64(arr[i]), float64(arr[i]*min)))

		ans = int(math.Max(float64(max), float64(ans)))
	}

	return ans
}
