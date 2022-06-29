package _01

import (
	"fmt"
	"math"
)

func getSpiralTransversal(arr [][]int) []int {
	var l, r, u, d = 0, len(arr[0]) - 1, 0, len(arr) - 1
	var result []int

	for d >= u && r >= l {
		for temp := l; temp <= r; temp++ {
			if arr[l][temp] != math.MaxInt {
				result = append(result, arr[l][temp])
				arr[l][temp] = math.MaxInt
			}
		}
		u++

		for temp := u; temp <= d; temp++ {
			if arr[temp][r] != math.MaxInt {
				result = append(result, arr[temp][r])
				arr[temp][r] = math.MaxInt
			}
		}
		r--

		for temp := r; temp >= l; temp-- {
			if arr[d][temp] != math.MaxInt {
				result = append(result, arr[d][temp])
				arr[d][temp] = math.MaxInt
			}
		}
		d--

		for temp := d; temp >= u; temp-- {
			if arr[temp][l] != math.MaxInt {
				result = append(result, arr[temp][l])
				arr[temp][l] = math.MaxInt
			}
		}
		l++
	}

	fmt.Println(result)
	return result
}
