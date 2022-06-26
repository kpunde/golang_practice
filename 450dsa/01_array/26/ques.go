package _26

func getMajorityElement(arr []int) []int {
	var track = make(map[int]int)
	var length = len(arr) / 3

	var res []int

	for _, item := range arr {
		track[item] += 1
		if track[item] == length+1 {
			res = append(res, item)
		}
	}

	return res
}

// Moore Voting Algo Time: O(n) Space: O(1)
// Ref: https://www.youtube.com/watch?v=yDbkQd9t2ig
// Ref: https://www.youtube.com/watch?v=X0G5jEcvroo
func getMajorityElementImproved(arr []int) []int {
	// There will be always mx of 2 nums with repeating greater than n/3
	num1, num2, c1, c2 := -1, -1, 0, 0
	for _, item := range arr {
		if item == num1 {
			c1++
		} else if item == num2 {
			c2++
		} else if c1 == 0 {
			num1 = item
			c1 = 1
		} else if c2 == 0 {
			num2 = item
			c2 = 1
		} else {
			c1--
			c2--
		}
	}

	c1, c2 = 0, 0
	for _, item := range arr {
		if item == num1 {
			c1++
		} else if item == num2 {
			c2++
		}
	}

	var res []int
	if c1 > len(arr)/3 {
		res = append(res, num1)
	}
	if c2 > len(arr)/3 {
		res = append(res, num2)
	}
	return res
}
