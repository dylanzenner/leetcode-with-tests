package permutation

func Permutate(nums []int) []int {
	res := make([]int, len(nums))

	for i, _ := range nums {
		res[i] = nums[nums[i]]
	}
	return res
}

