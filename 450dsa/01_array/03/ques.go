package _3

import (
	"math"
	"sort"
)

// O(nlogn)
func _kSmallestUsingSort(list []int, pos int) int {
	if pos-1 > len(list) {
		return math.MinInt
	}
	sort.Ints(list)
	return list[pos-1]
}

// Partition for quicksort/ quick select algo
func _partition(list []int, low, high int) ([]int, int) {
	pivot := list[high]
	i := low
	for j := low; j < high; j++ {
		if list[j] < pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}
	list[i], list[high] = list[high], list[i]
	return list, i
}

// Best case O(k), worst case O(n^2)
func _kSmallestUsingQuickSelect(list []int, pos int, start int, end int) int {
	for {
		lst, pivot := _partition(list, start, end)
		if pivot == pos {
			return lst[pos]
		} else if pos > pivot {
			list = lst
			start = pivot + 1
		} else if pos < pivot {
			list = lst
			end = pivot - 1
		}
	}
}

func getKSmallest(list []int, pos int) int {
	return _kSmallestUsingQuickSelect(list, pos-1, 0, len(list)-1)
}
