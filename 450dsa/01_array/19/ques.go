package _19

func getPairsCount(arr []int, k int) int {
	var mapping = make(map[int]int)
	count := 0
	for _, item := range arr {
		if val, ok := mapping[k-item]; ok {
			count += val
		}
		mapping[item]++
	}

	return count
}
