package sneakyNums

import "slices"

func SneakyNums(nums []int) []int {
	slices.Sort(nums)
	res := make([]int, 2)
	left, right := 0, 1
	pointer := 0

	for right < len(nums) {
		if nums[left] == nums[right] {
			res[pointer] = nums[left]
			right += 1
			left = right
			right += 1
			pointer += 1

		} else {
			left += 1
			right += 1
		}
	}

	return res
}