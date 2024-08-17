package maxones

func MaxOnes(nums []int) int {
	res := 0
	current := 0

	for _, val := range nums {
		if val == 1 {
			current += 1
		} else if val == 0 {
			if current > res {
				res = current
				current = 0
			}
			current =0 
		}
	}

	if current > res {
		res = current
	}

	return res
}