package minAverage

import (
	"slices"
	"math"
)

func MinAverage(nums []int) float64 {
	slices.Sort(nums)
	minimum := math.Inf(1)
	left, right := 0, len(nums) - 1

	for left < right {
		val := (float64(nums[left]) + float64(nums[right])) / 2.0
		if val < minimum {
			minimum = val
		}
		left += 1
		right -= 1
	}

	return minimum
}