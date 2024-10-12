package heightChecker

import "slices"

func HeightChecker(heights []int) int {
	heights_copy := slices.Clone(heights)
	slices.Sort(heights)
	res, pointer := 0, 0

	for pointer < len(heights) {
		if heights[pointer] != heights_copy[pointer] {
			res += 1
		}
		pointer += 1
	}

	return res
}