package runningsum

func RunningSum(nums []int) []int {
	res := []int{nums[0]}
	for _, val := range nums[1:]{
		res = append(res, res[len(res) - 1] + val)
	}
	return res
}