package _7

func getCyclicRotate(list []int) []int {
	var result []int
	result = append(result, list[len(list)-1])
	for i := 0; i <= len(list)-2; i++ {
		result = append(result, list[i])
	}
	return result
}
