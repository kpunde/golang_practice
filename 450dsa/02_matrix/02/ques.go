package _02

func getSearchMatrix(arr [][]int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if target >= arr[i][0] && target <= arr[i][len(arr[i])-1] {
			for j := 0; j < len(arr[i]); j++ {
				if arr[i][j] == target {
					return true
				}
			}
		}
	}

	return false
}
