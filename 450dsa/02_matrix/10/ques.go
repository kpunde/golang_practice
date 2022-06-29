package _10

func printCommonElements(arr [][]int) []int {
	var track []map[int]int
	var masterTrack = make(map[int]int)

	n := len(arr)

	var result []int

	for i := 0; i < len(arr); i++ {
		var temp = make(map[int]int)
		for j := 0; j < len(arr[i]); j++ {
			temp[arr[i][j]] = 1
		}
		track = append(track, temp)
	}

	for _, item := range track {
		for key, _ := range item {
			masterTrack[key] += 1
		}
	}

	for k, v := range masterTrack {
		if v == n {
			result = append(result, k)
		}
	}

	return result

}
