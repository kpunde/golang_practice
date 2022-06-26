package _34

func merge(arr []int) int {
	n := len(arr)
	i, j := 0, n-1
	count := 0

	for i < j {
		if arr[i] == arr[j] {
			i++
			j--
		} else if arr[i] > arr[j] {
			j--
			arr[j] += arr[j+1]
			count++
		} else {
			i++
			arr[i] += arr[i-1]
			count++
		}
	}

	return count
}
