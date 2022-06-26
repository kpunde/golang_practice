package _12

// Gap Algo Ref: https://www.youtube.com/watch?v=hVl2b3bLzBw
func mergeWithoutSpace(arr1 []int, arr2 []int) ([]int, []int) {
	maxIter := len(arr1) + len(arr2)

	for i := maxIter / 2; i > 0; i = i / 2 {
		for j := 0; j < maxIter; j += 1 {
			je := j + i

			if j < len(arr1) && je < len(arr1) {
				if arr1[j] > arr1[je] {
					arr1[j], arr1[je] = arr1[je], arr1[j]
				}
				continue
			}

			if j < len(arr1) && je >= len(arr1) {
				endPtr := je - len(arr1)
				if arr1[j] > arr2[endPtr] {
					arr1[j], arr2[endPtr] = arr2[endPtr], arr1[j]
				}
				continue
			}

			if j >= len(arr1) && je >= len(arr1) {
				endPtr := je - len(arr1)
				startPtr := j - len(arr1)

				if endPtr < len(arr2) && startPtr < len(arr2) {
					if arr2[startPtr] > arr2[endPtr] {
						arr2[startPtr], arr2[endPtr] = arr2[endPtr], arr2[startPtr]
					}
				}
				continue
			}

		}
	}

	return arr1, arr2
}
