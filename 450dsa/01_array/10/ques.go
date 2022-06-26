package _0

func getMinJumps(list []int) int {
	arrLen := len(list)
	jumpCount := 0

	i := 0
	for i < arrLen-1 {
		if list[i] == 0 {
			return -1
		}
		i = i + list[i]
		jumpCount++
	}

	return jumpCount
}
