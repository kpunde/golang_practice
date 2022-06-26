package _27

func getIsSubset(arr []int, arr02 []int) bool {
	var track = make(map[int]int)
	for _, item := range arr {
		track[item] = 1
	}

	for _, item := range arr02 {
		if _, ok := track[item]; !ok {
			return false
		}
	}

	return true
}
