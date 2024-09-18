package binarySearch

import "slices"

func BinarySearch(nums []int, target int) int {
	slices.Sort(nums)
	left, right := 0, len(nums) - 1

	for left <= right {
		m := (left + right) / 2
		if nums[m] == target {
			return m
		} else if target > nums[m] {
			left = m + 1
		} else if target < nums[m] {
			right = m - 1
		}
	}
	return -1
}