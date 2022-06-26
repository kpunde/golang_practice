package _6

func getUnion(list01 []int, list02 []int) int {
	var mapping = make(map[int]int)
	for _, item := range list01 {
		if _, ok := mapping[item]; !ok {
			mapping[item] = 1
		}
	}

	for _, item := range list02 {
		if _, ok := mapping[item]; !ok {
			mapping[item] = 1
		}
	}

	var result []int
	for k, _ := range mapping {
		result = append(result, k)
	}

	return len(result)
}
