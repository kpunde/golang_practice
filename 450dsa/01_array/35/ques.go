package _35

func findMedianSortedArrays(num1 []int, num2 []int) float64 {
	var resultantArr []int

	i, j := 0, 0

	for i < len(num1) && j < len(num2) {
		if num1[i] < num2[j] {
			resultantArr = append(resultantArr, num1[i])
			i++
		} else {
			resultantArr = append(resultantArr, num2[j])
			j++
		}
	}

	for i < len(num1) {
		resultantArr = append(resultantArr, num1[i])
		i++
	}
	for j < len(num2) {
		resultantArr = append(resultantArr, num2[j])
		j++
	}

	n := len(resultantArr)
	if n%2 == 0 {
		return float64(resultantArr[n/2-1]+resultantArr[n/2]) / 2
	} else {
		return float64(resultantArr[n/2])
	}

}
