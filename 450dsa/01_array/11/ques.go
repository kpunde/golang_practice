package _11

func findDuplicateN1(list []int) int {
	var mapping = make(map[int]int)
	for _, item := range list {
		if _, ok := mapping[item]; ok {
			return item
		}
		mapping[item] = 1
	}
	return 0
}
