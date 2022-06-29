package _08

func getRotate90Clockwise(arr [][]int) [][]int {
	var temp = make([][]int, len(arr))
	for i := range temp {
		temp[i] = make([]int, len(arr[0]))
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			temp[j][len(arr[i])-1-i] = arr[i][j]
		}
	}
	return temp
}
