package maxProd

import "slices"

func MaxProd(nums []int) int {
	slices.Sort(nums)
	length := len(nums)
	return (nums[length - 1] - 1) * (nums[length - 2] - 1)
}