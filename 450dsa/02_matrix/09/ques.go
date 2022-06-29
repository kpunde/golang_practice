package _09

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

// Quick Selection : Best case O(k), worst case O(n^2)
func _kSmallestUsingQuickSelect(arr []int, pos int, start int, end int) int {
	for {
		lst, pivot := _partition(arr, start, end)
		if pivot == pos {
			return lst[pivot]
		} else if pivot > pos {
			end = pivot - 1
		} else {
			start = pivot + 1
		}
	}
}

func getKSmallestArray(list []int, pos int) int {
	return _kSmallestUsingQuickSelect(list, pos-1, 0, len(list)-1)
}

func getKthSmallest(arr [][]int, k int) int {
	var temp []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			temp = append(temp, arr[i][j])
		}
	}

	return getKSmallestArray(temp, k)
}
