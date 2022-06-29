package _05

import "sort"

func getSortedMatrix(arr [][]int) [][]int {
	var temp []int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			temp = append(temp, arr[i][j])
		}
	}

	sort.Ints(temp)

	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = temp[count]
			count++
		}
	}
	return arr
}
