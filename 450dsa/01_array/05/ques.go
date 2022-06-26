package _5

func moveNegativeToEnd(list []int) []int {
	var mapping = make(map[int][]int)
	for _, item := range list {
		if item < 0 {
			if value, ok := mapping[-1]; ok {
				mapping[-1] = append(value, item)
			} else {
				mapping[-1] = []int{item}
			}
		} else {
			if value, ok := mapping[1]; ok {
				mapping[1] = append(value, item)
			} else {
				mapping[1] = []int{item}
			}
		}
	}

	var result []int
	result = append(mapping[1], mapping[-1]...)
	return result
}
