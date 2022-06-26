package _20

func getCommonElements(arr01 []int, arr02 []int, arr03 []int) []int {
	var result []int
	var mapping = make(map[int]int)
	i, j, k := 0, 0, 0
	for i < len(arr01) && j < len(arr02) && k < len(arr03) {
		if arr01[i] == arr02[j] && arr02[j] == arr03[k] {
			if _, ok := mapping[arr01[i]]; !ok {
				mapping[arr01[i]] = 1
			}
			i += 1
			j += 1
			k += 1
			continue
		}
		if i < len(arr01) && j < len(arr02) && k < len(arr03) {
			if arr01[i] < arr02[j] || arr01[i] < arr03[k] {
				i++
			}
		}

		if i < len(arr01) && j < len(arr02) && k < len(arr03) {
			if arr02[j] < arr03[k] || arr02[j] < arr01[i] {
				j++
			}
		}

		if i < len(arr01) && j < len(arr02) && k < len(arr03) {
			if arr03[k] < arr02[j] || arr03[k] < arr01[i] {
				k++
			}
		}
	}

	for k, _ := range mapping {
		result = append(result, k)
	}
	return result
}
