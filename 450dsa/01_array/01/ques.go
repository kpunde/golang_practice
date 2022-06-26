package _1

func reverseList(list []int) []int {
	var finalList []int
	for i := len(list) - 1; i >= 0; i-- {
		finalList = append(finalList, list[i])
	}
	return finalList
}
