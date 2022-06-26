package permutation

func recurPermute(arr *[]int, ans *[][]int, index int) {
	//base case
	if index >= len(*arr) {
		var temp []int
		for _, item := range *arr {
			temp = append(temp, item)
		}
		*ans = append(*ans, temp)
		return
	}

	for i := index; i < len(*arr); i++ {
		(*arr)[index], (*arr)[i] = (*arr)[i], (*arr)[index]
		recurPermute(arr, ans, index+1)
		(*arr)[index], (*arr)[i] = (*arr)[i], (*arr)[index]
	}
}

func permute(arr []int) [][]int {
	var ans [][]int
	var index = 0
	recurPermute(&arr, &ans, index)
	return ans
}
