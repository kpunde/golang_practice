package _03

import "sort"

func getMedian(arr [][]int) int {
	var temp []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			temp = append(temp, arr[i][j])
		}
	}
	sort.Ints(temp)
	if len(temp)%2 == 0 {
		return (temp[len(temp)/2] + temp[len(temp)/2+1]) / 2
	} else {
		return temp[len(temp)/2]
	}
}
