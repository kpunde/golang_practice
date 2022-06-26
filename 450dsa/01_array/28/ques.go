package _28

func _removeElem(pos int, arr []int) []int {
	var res []int
	res = append(arr[:pos], arr[pos+1:]...)
	return res
}

func find3Numbers(arr []int, sum int) bool {

	for index, item := range arr {

		k := sum - item
		newArr := _removeElem(index, arr)
		mapping := make(map[int]int)
		for _, inItem := range newArr {
			if _, ok := mapping[inItem]; ok {
				return true
			}
			mapping[k-inItem] = 1
		}

	}
	return false
}
