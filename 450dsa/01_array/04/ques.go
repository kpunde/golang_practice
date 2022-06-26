package _4

func sort012(list []int) []int {
	var mapping = make(map[int][]int)
	for _, item := range list {
		switch item {
		case 0:
			if val, ok := mapping[item]; ok {
				mapping[0] = append(val, item)
				break
			}

			mapping[0] = []int{item}
			break
		case 1:
			if val, ok := mapping[item]; ok {
				mapping[item] = append(val, item)
				break
			}

			mapping[item] = []int{item}
			break
		case 2:
			if val, ok := mapping[item]; ok {
				mapping[item] = append(val, item)
				break
			}

			mapping[item] = []int{item}
			break
		}
	}

	var result []int

	result = append(mapping[0], mapping[1]...)
	result = append(result, mapping[2]...)

	return result
}
