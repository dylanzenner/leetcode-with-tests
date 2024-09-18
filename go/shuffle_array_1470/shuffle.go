package shuffle


func Shuffle(nums []int, n int) []int {
	left, right := 0, len(nums) - n
	pointer := 0
	res := make([]int, len(nums))

	for right < len(nums) {
		res[pointer] = nums[left]
		pointer += 1
		res[pointer] = nums[right]
		pointer += 1

		left += 1
		right += 1
	}

	return res
}