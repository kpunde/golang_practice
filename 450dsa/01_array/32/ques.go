package _32

func threeWayPartition(arr []int, a int, b int) []int {
	toSwap, i := -1, 0
	for ; i < len(arr); i++ {
		if arr[i] < a {
			toSwap++
			arr[toSwap], arr[i] = arr[i], arr[toSwap]
		}
	}

	for i = toSwap + 1; i < len(arr); i++ {
		if arr[i] <= b {
			toSwap++
			arr[toSwap], arr[i] = arr[i], arr[toSwap]
		}
	}

	for i := toSwap + 1; i < len(arr); i++ {
		if arr[i] > b {
			toSwap++
			arr[toSwap], arr[i] = arr[i], arr[toSwap]
		}
	}

	return arr
}
