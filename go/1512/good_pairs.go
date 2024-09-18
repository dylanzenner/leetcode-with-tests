package goodPairs

import "slices"

func GoodPairs(nums []int) int {
	slices.Sort(nums)
	count := 0
	left, right := 0, 1

	for left < len(nums) && right < len(nums) {
		if nums[left] == nums[right] && right == len(nums) - 1 {
			count += 1
			left += 1
			right = left + 1
		} else if nums[left] == nums[right] {
			count += 1
			right += 1
		} else if nums[left] != nums[right] {
			left += 1
			right = left + 1
		}
	}

	return count

}