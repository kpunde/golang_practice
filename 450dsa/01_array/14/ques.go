package _14

func getMerge(arr [][]int) [][]int {
	var result [][]int
	start, end := arr[0][0], arr[0][1]

	for index, item := range arr {
		if index == 0 {
			continue
		}

		if item[0] <= end {
			end = item[1]
		} else {
			result = append(result, []int{start, end})
			start = item[0]
			end = item[1]
		}
	}
	result = append(result, []int{start, end})

	return result
}
