package _33

import "math"

// sliding window
func minSwap(arr []int, k int) int {
	fav, nonFav := 0, 0

	for _, item := range arr {
		if item <= k {
			fav++
		}
	}

	for i := 0; i < fav; i++ {
		if arr[i] > k {
			nonFav++
		}
	}
	l, r := 0, fav-1
	result := math.MaxInt

	for r < len(arr) {
		result = int(math.Min(float64(result), float64(nonFav)))
		r++
		if r < len(arr) && arr[r] > k {
			nonFav++
		}
		if r < len(arr) && arr[l] > k {
			nonFav--
		}
	}

	if result == math.MaxInt {
		return 0
	}

	return result
}
