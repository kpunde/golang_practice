package _24

func _noToArray(num int) []int {
	q, r := 0, 0
	var arr []int
	for num != 0 {
		q = num / 10
		r = num % 10
		arr = append([]int{r}, arr...)
		num = q
	}
	return arr
}

func _getArrNumProduct(arr []int, num int) []int {
	handler, carry := 0, 0
	var result []int
	for i := len(arr) - 1; i >= 0; i-- {
		handler = num*arr[i] + carry
		carry = handler / 10
		handler = handler % 10
		result = append([]int{handler}, result...)
	}
	result = append(_noToArray(carry), result...)
	return result
}

func getFactorial(num int) []int {
	if num <= 1 {
		return []int{1}
	}
	return _getArrNumProduct(getFactorial(num-1), num)
}
