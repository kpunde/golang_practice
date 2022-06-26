package _29

import "math"

// explanation: https://www.youtube.com/watch?v=C8UjlJZsHBw
func trappingWater(arr []int) int {
	var lMax []int
	var rMax []int

	lM, rM := 0, 0
	vol := 0

	for i, j := 0, len(arr)-1; i < len(arr) && j >= 0; i, j = i+1, j-1 {
		if arr[i] > lM {
			lM = arr[i]
		}

		if arr[j] > rM {
			rM = arr[j]
		}

		lMax = append(lMax, lM)
		rMax = append(rMax, rM)
	}
	for i := 0; i < len(arr); i++ {
		vol += int(math.Min(float64(lMax[i]), float64(rMax[i]))) - arr[i]
	}

	return vol
}
