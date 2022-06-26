package _22

// Move to the right of an array while adding to a sum, if there is a repeating of any digit, that means there is 0 present
func subArrayExists(arr []int) bool {
	var mapping = make(map[int]int)
	sum := 0
	for _, item := range arr {
		sum += item
		if _, ok := mapping[sum]; ok {
			return true
		}
		mapping[sum] = 1
	}
	return false
}
