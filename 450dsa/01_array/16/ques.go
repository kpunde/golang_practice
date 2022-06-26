package _16

import "sort"

func getInversionCount(arr []int) int {
	var count int
	var ref []int

	for _, item := range arr {
		ref = append(ref, item)
	}
	sort.Ints(ref)

	for index, item := range ref {
		if item != arr[index] {
			count++
		}
	}

	if count > 0 {
		return count - 1
	}
	return count
}
