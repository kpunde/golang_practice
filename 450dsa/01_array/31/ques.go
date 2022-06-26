package _31

import "math"

func minSubArrayLen(target int, arr []int) int {
	count := math.MaxInt
	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		lCount := 1
		if sum >= target {
			count = int(math.Min(float64(count), float64(lCount)))
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			lCount++
			if sum >= target {
				count = int(math.Min(float64(count), float64(lCount)))
			}
		}
	}
	if count == math.MaxInt {
		return 0
	}
	return count
}

// sliding window approach
func minSubArrayLenImproved(target int, arr []int) int {
	st, en, sum, n, count := 0, 0, arr[0], len(arr), math.MaxInt

	for st <= en && en < n {
		if sum < target {
			en++
			if en < n {
				sum += arr[en]
			}
		} else {
			count = int(math.Min(float64(count), float64(en-st+1)))
			sum -= arr[st]
			st++
		}
	}

	if count == math.MaxInt {
		return 0
	}
	return count
}
