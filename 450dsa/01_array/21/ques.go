package _21

func getRearrange(arr []int) []int {
	var positive, negative []int
	for _, item := range arr {
		if item < 0 {
			negative = append(negative, item)
			continue
		}
		positive = append(positive, item)
	}

	i, j, ptr := 0, 0, 0

	for i < len(positive) && j < len(negative) {
		if ptr%2 == 0 {
			arr[ptr] = positive[i]
			i++
		} else {
			arr[ptr] = negative[j]
			j++
		}
		ptr++
	}

	for i < len(positive) && ptr < len(arr) {
		arr[ptr] = positive[i]
		ptr++
		i++
	}

	for j < len(negative) && ptr < len(arr) {
		arr[ptr] = negative[j]
		ptr++
		j++
	}

	return arr
}
